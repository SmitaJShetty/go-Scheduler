package taskqueue

import "goscheduler/go-Scheduler/internal/model"

//TaskQueue queue that contains tasks
type TaskQueue interface {
	StoreTasks(task *model.Task)
	FetchNext() (*model.Task, error)
}
