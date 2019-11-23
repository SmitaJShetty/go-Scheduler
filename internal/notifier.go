package internal

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/smitajshetty/go-scheduler/internal/model"
	"github.com/smitajshetty/go-scheduler/internal/worker"
)

//Notifier schedules
type Notifier struct {
	Client *redis.Client
}

//NewNotifier returns a redisqueue
func NewNotifier() *Notifier {
	return &Notifier{
		Client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	}
}

/*
	checks on processing queueu
	if tasks exist with extended ttl, then create failure log
*/
//CheckTaskProcessingStatus set up task execution times
func (n *Notifier) CheckTaskProcessingStatus() error {
	result, readErr := n.Client.XRead(&redis.XReadArgs{
		Streams: []string{worker.ProcessingQueueName, "0"},
		Block:   100 * time.Millisecond,
	}).Result()
	if readErr != nil {
		return readErr
	}

	if result == nil {
		return fmt.Errorf("read result was empty")
	}

	t, getErr := n.getTaskFromResult(result)
	if getErr != nil {
		return getErr
	}

	if t == nil {
		return fmt.Errorf("invalid task")
	}

	if t.Status == model.Inactive {
		return fmt.Errorf("cleanup of task %v was not performed", t)
	}

	if time.Add(t.StartTime+t.TTLInSeconds) <= time.Now() { // work had to be completed by now, but isnt
		return util.LogErrorTrail(fmt.Sprintf("Job was to be done by now, but isnt, details of task as follows: %v", t))
	}

	return nil
}

func (n *Notifier) getTaskFromResult(result string) (*model.Task, error) {
	if result == "" {
		return nil, fmt.Errorf("invalid result string")
	}

	var t model.Task
	marshalErr := json.Unmarshal([]byte(result), &t)
	if marshalErr != nil {
		return nil, fmt.Errorf("error while unmarshalling result : %s, %v", result, marshallErr)
	}

	return &t, nil
}
