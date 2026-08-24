package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/jakkayy/devHub/services/backend-api/internal/domain"
	"github.com/jakkayy/devHub/services/backend-api/internal/repository"
)

type ContextLinkerUsecase struct {
	linkerRepo   domain.ContextLinkerRepositoryInterface
	contractRepo domain.APIContractRepositoryInterface
	cacheRepo    *repository.CacheRepository
}

func NewContextLinkerUsecase(
	linkerRepo domain.ContextLinkerRepositoryInterface,
	contractRepo domain.APIContractRepositoryInterface,
	cacheRepo *repository.CacheRepository,
) *ContextLinkerUsecase {
	return &ContextLinkerUsecase{
		linkerRepo:   linkerRepo,
		contractRepo: contractRepo,
		cacheRepo:    cacheRepo,
	}
}

// GenerateVSCodeDeepLink generates custom vscode:// URI protocol
func (u *ContextLinkerUsecase) GenerateVSCodeDeepLink(filePath string, lineNumber int) string {
	if lineNumber > 0 {
		return fmt.Sprintf("vscode://file/%s:%d", filePath, lineNumber)
	}
	return fmt.Sprintf("vscode://file/%s", filePath)
}

// CreateContextLink binds Task, File Path, and API Contract Spec together
func (u *ContextLinkerUsecase) CreateContextLink(ctx context.Context, dto domain.CreateLinkDTO) (*domain.ContextLinkDetailsDTO, error) {
	linkID := fmt.Sprintf("link-%d", time.Now().UnixNano())

	link := &domain.ContextLink{
		ID:            linkID,
		TaskID:        dto.TaskID,
		APIContractID: dto.APIContractID,
		FilePath:      dto.FilePath,
		GitBranch:     dto.GitBranch,
		CreatedAt:     time.Now(),
	}

	if err := u.linkerRepo.CreateLink(ctx, link); err != nil {
		return nil, fmt.Errorf("failed to save context link: %w", err)
	}

	deepLink := u.GenerateVSCodeDeepLink(dto.FilePath, dto.LineNumber)

	var contract *domain.APIContract
	if dto.APIContractID != "" {
		contract, _ = u.contractRepo.GetContractByID(ctx, dto.APIContractID)
	}

	return &domain.ContextLinkDetailsDTO{
		ID:            link.ID,
		TaskID:        link.TaskID,
		APIContractID: link.APIContractID,
		APIContract:   contract,
		FilePath:      link.FilePath,
		GitBranch:     link.GitBranch,
		VSCodeURI:     deepLink,
	}, nil
}

// GetLinksForTask retrieves all context links associated with a Task ID
func (u *ContextLinkerUsecase) GetLinksForTask(ctx context.Context, taskID string) ([]domain.ContextLinkDetailsDTO, error) {
	links, err := u.linkerRepo.GetLinksByTaskID(ctx, taskID)
	if err != nil {
		return nil, err
	}

	var result []domain.ContextLinkDetailsDTO
	for _, l := range links {
		deepLink := u.GenerateVSCodeDeepLink(l.FilePath, 0)
		var contract *domain.APIContract
		if l.APIContractID != "" {
			contract, _ = u.contractRepo.GetContractByID(ctx, l.APIContractID)
		}

		result = append(result, domain.ContextLinkDetailsDTO{
			ID:            l.ID,
			TaskID:        l.TaskID,
			APIContractID: l.APIContractID,
			APIContract:   contract,
			FilePath:      l.FilePath,
			GitBranch:     l.GitBranch,
			VSCodeURI:     deepLink,
		})
	}

	return result, nil
}
