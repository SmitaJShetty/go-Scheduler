package worker

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/smitajshetty/go-scheduler/internal/model"
)

//ProcessingQueueName processing queue
const ProcessingQueueName string = "ProcessingQueue"

//TaskQueueName name of the task queue
const TaskQueueName string = "TaskQueue"

//NewWorker returns a redisqueue
func NewWorker() *Worker {
	return &Worker{
		Client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	}
}

/*
	add to processing queue
	scheduler checks tasks on the opposite end
	if task's TTL + starttime <= currenttime, then task is removed from queue, placed under failed queue
	if task successfully processed, worker logs in success log
*/

//Worker queue to store tasks in redis
type Worker struct {
	Client *redis.Client
}

//AddTask stores tasks
func (w *Worker) AddTask(task *model.Task) error {
	/* store tasks -- stores task in redis queue*/
	/* if queue does not exist, ceeate it and add it*/
	if w.Client == nil {
		return fmt.Errorf("client empty")
	}

	pushErr := w.Client.LPush(TaskQueueName, task).Err()
	if pushErr != nil {
		return pushErr
	}
	return nil
}

//RemoveTask remove from task queue and moves a copy into processing queue
func (w *Worker) RemoveTask() (*model.Task, error) {
	//remove from queue and store in processing queue
	taskStr, popErr := w.Client.RPopLPush(TaskQueueName, ProcessingQueueName).Result()
	if popErr != nil {
		return nil, popErr
	}

	var t model.Task
	jsonErr := json.Unmarshal([]byte(taskStr), &t)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &t, nil
}
