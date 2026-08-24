package repository

import (
	"context"

	"github.com/jakkayy/devHub/services/backend-api/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type APIContractRepository struct {
	db *gorm.DB
}

func NewAPIContractRepository(db *gorm.DB) *APIContractRepository {
	return &APIContractRepository{db: db}
}

// UpsertContracts saves or updates API contracts in PostgreSQL
func (r *APIContractRepository) UpsertContracts(ctx context.Context, contracts []domain.APIContract) error {
	if len(contracts) == 0 {
		return nil
	}
	return r.db.WithContext(ctx).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&contracts).Error
}

// GetAllContracts retrieves all API contracts from PostgreSQL
func (r *APIContractRepository) GetAllContracts(ctx context.Context) ([]domain.APIContract, error) {
	var contracts []domain.APIContract
	err := r.db.WithContext(ctx).Find(&contracts).Error
	return contracts, err
}

// GetContractByID retrieves a single API contract spec by ID
func (r *APIContractRepository) GetContractByID(ctx context.Context, id string) (*domain.APIContract, error) {
	var contract domain.APIContract
	err := r.db.WithContext(ctx).First(&contract, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &contract, nil
}
