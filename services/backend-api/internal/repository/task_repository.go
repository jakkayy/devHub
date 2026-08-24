package repository

import (
	"context"

	"github.com/jakkayy/devHub/services/backend-api/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

// UpsertTasks saves or updates tasks in PostgreSQL based on Primary Key (ID)
func (r *TaskRepository) UpsertTasks(ctx context.Context, tasks []domain.Task) error {
	if len(tasks) == 0 {
		return nil
	}
	return r.db.WithContext(ctx).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&tasks).Error
}

// GetAllTasks retrieves all tasks from PostgreSQL
func (r *TaskRepository) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	var tasks []domain.Task
	err := r.db.WithContext(ctx).Find(&tasks).Error
	return tasks, err
}
