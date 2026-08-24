package usecase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jakkayy/devHub/services/backend-api/internal/domain"
	"github.com/jakkayy/devHub/services/backend-api/internal/repository"
)

type APIContractUsecase struct {
	parser       domain.APIContractParserInterface
	contractRepo domain.APIContractRepositoryInterface
	cacheRepo    *repository.CacheRepository
}

func NewAPIContractUsecase(
	parser domain.APIContractParserInterface,
	contractRepo domain.APIContractRepositoryInterface,
	cacheRepo *repository.CacheRepository,
) *APIContractUsecase {
	return &APIContractUsecase{
		parser:       parser,
		contractRepo: contractRepo,
		cacheRepo:    cacheRepo,
	}
}

// SyncAndGetContracts parses raw spec JSON, caches specs in Redis, and upserts them to PostgreSQL
func (u *APIContractUsecase) SyncAndGetContracts(ctx context.Context, rawSpec []byte) ([]domain.APIContract, error) {
	cacheKey := "api_contracts:all"

	// 1. Try Redis Cache first if no new raw spec provided
	if len(rawSpec) == 0 {
		var cachedContracts []domain.APIContract
		if err := u.cacheRepo.Get(ctx, cacheKey, &cachedContracts); err == nil && len(cachedContracts) > 0 {
			return cachedContracts, nil
		}
	}

	// 2. Parse OpenAPI Specs
	dtos, err := u.parser.ParseSpec(ctx, rawSpec)
	if err != nil {
		log.Printf("[APIContractUsecase] Parse error: %v. Falling back to DB.", err)
		return u.contractRepo.GetAllContracts(ctx)
	}

	// 3. Convert DTOs to Domain Entities
	var contracts []domain.APIContract
	for _, dto := range dtos {
		contracts = append(contracts, domain.APIContract{
			ID:          dto.ID,
			Method:      dto.Method,
			Path:        dto.Path,
			Summary:     dto.Summary,
			SpecVersion: dto.SpecVersion,
		})
	}

	// 4. Save to PostgreSQL Database
	if err := u.contractRepo.UpsertContracts(ctx, contracts); err != nil {
		log.Printf("[APIContractUsecase] Failed to upsert contracts: %v", err)
	}

	// 5. Store in Redis Cache (TTL 10 minutes)
	_ = u.cacheRepo.Set(ctx, cacheKey, contracts, 10*time.Minute)

	return contracts, nil
}

// GetContractByID fetches a specific contract spec by ID
func (u *APIContractUsecase) GetContractByID(ctx context.Context, id string) (*domain.APIContract, error) {
	cacheKey := fmt.Sprintf("api_contracts:id:%s", id)

	var cachedContract domain.APIContract
	if err := u.cacheRepo.Get(ctx, cacheKey, &cachedContract); err == nil && cachedContract.ID != "" {
		return &cachedContract, nil
	}

	contract, err := u.contractRepo.GetContractByID(ctx, id)
	if err != nil {
		return nil, err
	}

	_ = u.cacheRepo.Set(ctx, cacheKey, contract, 10*time.Minute)
	return contract, nil
}
