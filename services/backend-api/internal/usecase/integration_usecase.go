package usecase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jakkayy/devHub/services/backend-api/internal/domain"
	"github.com/jakkayy/devHub/services/backend-api/internal/repository"
)

type IntegrationUsecase struct {
	sheetsClient domain.SheetsClientInterface
	githubClient domain.GitHubClientInterface
	taskRepo     *repository.TaskRepository
	cacheRepo    *repository.CacheRepository
}

func NewIntegrationUsecase(
	sheetsClient domain.SheetsClientInterface,
	githubClient domain.GitHubClientInterface,
	taskRepo *repository.TaskRepository,
	cacheRepo *repository.CacheRepository,
) *IntegrationUsecase {
	return &IntegrationUsecase{
		sheetsClient: sheetsClient,
		githubClient: githubClient,
		taskRepo:     taskRepo,
		cacheRepo:    cacheRepo,
	}
}

// SyncAndGetTasks syncs tasks from Google Sheets, caches them in Redis, and stores them in Postgres
func (u *IntegrationUsecase) SyncAndGetTasks(ctx context.Context, spreadsheetID string) ([]domain.Task, error) {
	cacheKey := fmt.Sprintf("tasks:sheet:%s", spreadsheetID)

	// 1. Try Redis Cache first
	var cachedTasks []domain.Task
	if err := u.cacheRepo.Get(ctx, cacheKey, &cachedTasks); err == nil && len(cachedTasks) > 0 {
		return cachedTasks, nil
	}

	// 2. Fetch from Google Sheets API
	sheetDTOs, err := u.sheetsClient.FetchTasks(ctx, spreadsheetID, "Sheet1!A1:Z100")
	if err != nil {
		log.Printf("[Usecase] Sheets API error: %v. Falling back to DB.", err)
		return u.taskRepo.GetAllTasks(ctx)
	}

	// 3. Convert DTOs to Domain Tasks
	var tasks []domain.Task
	now := time.Now()
	for _, dto := range sheetDTOs {
		tasks = append(tasks, domain.Task{
			ID:          dto.TaskID,
			Title:       dto.Title,
			Assignee:    dto.Assignee,
			Status:      dto.Status,
			Source:      "google_sheets",
			ExternalURL: dto.ExternalURL,
			UpdatedAt:   now,
		})
	}

	// 4. Save to PostgreSQL async/sync
	if err := u.taskRepo.UpsertTasks(ctx, tasks); err != nil {
		log.Printf("[Usecase] Failed to upsert tasks to Postgres: %v", err)
	}

	// 5. Store in Redis Cache (TTL 3 minutes)
	_ = u.cacheRepo.Set(ctx, cacheKey, tasks, 3*time.Minute)

	return tasks, nil
}

// GetGitHubPullRequests fetches PRs with Redis caching (TTL 2 minutes)
func (u *IntegrationUsecase) GetGitHubPullRequests(ctx context.Context, owner string, repo string) ([]domain.GitHubPullRequestDTO, error) {
	cacheKey := fmt.Sprintf("github:pulls:%s:%s", owner, repo)

	var cachedPulls []domain.GitHubPullRequestDTO
	if err := u.cacheRepo.Get(ctx, cacheKey, &cachedPulls); err == nil && len(cachedPulls) > 0 {
		return cachedPulls, nil
	}

	pulls, err := u.githubClient.FetchPullRequests(ctx, owner, repo)
	if err != nil {
		return nil, err
	}

	_ = u.cacheRepo.Set(ctx, cacheKey, pulls, 2*time.Minute)
	return pulls, nil
}
