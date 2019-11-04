package taskqueue

import "goscheduler/go-Scheduler/internal/model"

//RedisQueue queue to store tasks in redis
type RedisQueue struct {
}

//StoreTasks stores tasks
func (rds *RedisQueue) StoreTasks(task *model.Task) error {
	/* store tasks -- stores task in redis queue*/
}

//FetchNext fetch next task
func (rds *RedisQueue) FetchNext() (*model.Task, error) {
	/* fetch next fetches the next task in queue waitig for processing*/
}
