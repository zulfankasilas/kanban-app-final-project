package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks(ctx context.Context, id int) ([]entity.Task, error)
	StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error)
	GetTaskByID(ctx context.Context, id int) (entity.Task, error)
	GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error)
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id int) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (r *taskRepository) GetTasks(ctx context.Context, id int) ([]entity.Task, error) {
	var result []entity.Task
	r.db.WithContext(ctx)

	err := r.db.Where("user_id = ?", id).Find(&result).Error

	if err != nil {
		return []entity.Task{}, err
	}

	return result, nil
}

func (r *taskRepository) StoreTask(ctx context.Context, task *entity.Task) (taskId int, err error) {
	r.db.WithContext(ctx)

	if err := r.db.Create(&task).Error; err != nil {
		return 0, err
	}

	return task.ID, nil
}

func (r *taskRepository) GetTaskByID(ctx context.Context, id int) (entity.Task, error) {
	var result entity.Task
	r.db.WithContext(ctx)

	err := r.db.Where("id = ?", id).First(&result).Error

	if err != nil {
		return entity.Task{}, err
	}

	return result, nil
}

func (r *taskRepository) GetTasksByCategoryID(ctx context.Context, catId int) ([]entity.Task, error) {
	var result []entity.Task
	r.db.WithContext(ctx)

	err := r.db.Where("category_id = ?", catId).Find(&result).Error

	if err != nil {
		return []entity.Task{}, err
	}

	return result, nil
}

func (r *taskRepository) UpdateTask(ctx context.Context, task *entity.Task) error {
	r.db.WithContext(ctx)

	if err := r.db.Model(&entity.Task{}).Where("id = ?", task.ID).Updates(&task).Error; err != nil {
		return err
	}

	return nil
}

func (r *taskRepository) DeleteTask(ctx context.Context, id int) error {
	r.db.WithContext(ctx)

	if err := r.db.Where("id = ?", id).Delete(&entity.Task{}).Error; err != nil {
		return err
	}

	return nil
}
