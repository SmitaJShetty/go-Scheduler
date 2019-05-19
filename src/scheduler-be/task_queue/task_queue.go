package task_queue

import "scheduler/go-Scheduler/src/scheduler-be/model"

//TaskQueue queue that contains tasks
type TaskQueue interface {
	StoreTasks(task *model.Task)
}
