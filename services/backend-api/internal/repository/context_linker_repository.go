package repository

import (
	"context"

	"github.com/jakkayy/devHub/services/backend-api/internal/domain"
	"gorm.io/gorm"
)

type ContextLinkerRepository struct {
	db *gorm.DB
}

func NewContextLinkerRepository(db *gorm.DB) *ContextLinkerRepository {
	return &ContextLinkerRepository{db: db}
}

// CreateLink saves a new context link mapping in PostgreSQL
func (r *ContextLinkerRepository) CreateLink(ctx context.Context, link *domain.ContextLink) error {
	return r.db.WithContext(ctx).Create(link).Error
}

// GetLinksByTaskID retrieves all links associated with a specific Task ID
func (r *ContextLinkerRepository) GetLinksByTaskID(ctx context.Context, taskID string) ([]domain.ContextLink, error) {
	var links []domain.ContextLink
	err := r.db.WithContext(ctx).Where("task_id = ?", taskID).Find(&links).Error
	return links, err
}

// GetLinkByID retrieves a single context link by ID
func (r *ContextLinkerRepository) GetLinkByID(ctx context.Context, id string) (*domain.ContextLink, error) {
	var link domain.ContextLink
	err := r.db.WithContext(ctx).First(&link, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &link, nil
}
