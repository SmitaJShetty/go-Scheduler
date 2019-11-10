package worker

//WorkerGenerator construct for worker
type WorkerGenerator struct {
}

//WorkerGenerator schedules a tasks
func (w *WorkerGenerator) ScheduleTask() {
	//create workers
	go generateCreateWorker()
	//pick up events and create tasks
}

func generateCreateWorker() {
	//pick a worker from queue

	//execute task

}
