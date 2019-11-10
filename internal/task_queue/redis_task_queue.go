package taskqueue

import "scheduler/go-Scheduler/src/scheduler-be/model"

//NewRedisQueue returns a redisqueue
func NewRedisQueue() *RedisQueue {
	return &RedisQueue{}
}

//RedisQueue queue to store tasks in redis
type RedisQueue struct {
}

//StoreTask stores tasks
func (rds *RedisQueue) StoreTask(task *model.Task) error {
	/* store tasks -- stores task in redis queue*/

}

//FetchNext fetch next task
func (rds *RedisQueue) FetchNext() (*model.Task, error) {
	/* fetch next fetches the next task in queue waitig for processing*/

}
