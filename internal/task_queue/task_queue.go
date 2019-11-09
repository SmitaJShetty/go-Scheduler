package taskqueue

import "scheduler/go-Scheduler/src/scheduler-be/model"

//TaskQueue queue that contains tasks
type TaskQueue interface {
	StoreTask(task *model.Task)
	FetchNext() (*model.Task, error)
}
