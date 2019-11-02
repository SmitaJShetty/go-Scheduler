package scheduler

import (
	"jwt/src/backend/repo"
)

//Scheduler schedules
type Scheduler struct {
	eventtasker
}

//ScheduleTasks set up task execution times
func (s *Scheduler) ScheduleTasks() {
	/*
		if events exist unattended, make tasks based on set up times
	*/

	
	

	 
	
}

func (s *Scheduler) pickEvents() {
	//
}

func (s *Scheduler) createTasks() {
	//pick events

}
