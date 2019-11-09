package repo

import (
	"fmt"
	"scheduler/go-Scheduler/internal/model"

	"github.com/jinzhu/gorm"
)

//TaskRepo construct for task
type TaskRepo struct {
	db *gorm.DB
}

//NewTaskRepo returns a new task repo
func NewTaskRepo() *TaskRepo {
	return &TaskRepo{
		db: &gorm.DB{},
	}
}

//Get gets id string
func (t *TaskRepo) Get(id string) (*model.Task, error) {
	if id == "" {
		return nil, fmt.Errorf("id is empty")
	}

	var t1 model.Task
	db := t.db.Where(model.Task{ID: id}).Find(&t1)
	if db.Error != nil {
		return nil, db.Error
	}

	if db.RecordNotFound() {
		return nil, nil
	}

	return &t1, nil
}

//Update updates event
func (t *TaskRepo) Update(task *model.Task) (*model.Task, error) {
	if task == nil {
		return nil, fmt.Errorf("id is empty")
	}

	db := t.db.Save(&task)
	if db.Error == nil {
		return nil, db.Error
	}

	if db.RecordNotFound() {
		return nil, nil
	}
	return task, nil
}

//Delete deletes a task based on id
func (t *TaskRepo) Delete(id string) error {
	if id == "" {
		return fmt.Errorf("id is empty")
	}

	task := &model.Task{
		ID: id,
	}

	db := t.db.Delete(task)
	if db.Error != nil {
		return db.Error
	}

	if db.RowsAffected == 0 {
		return nil
	}
	return nil
}

//Create creates a new task
func (t *TaskRepo) Create(task *model.Task) (*model.Task, error) {
	if task == nil {
		return nil, fmt.Errorf("task is empty")
	}
	db := t.db.Create(task)
	if db.Error != nil {
		return nil, db.Error
	}

	if db.RowsAffected == 0 {
		return nil, nil
	}
	return task, nil
}
