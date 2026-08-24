package domain

import "context"

// CreateLinkDTO represents payload for binding Task, File Path, and API Spec together
type CreateLinkDTO struct {
	TaskID        string `json:"task_id" binding:"required"`
	APIContractID string `json:"api_contract_id"`
	FilePath      string `json:"file_path" binding:"required"`
	LineNumber    int    `json:"line_number"`
	GitBranch     string `json:"git_branch"`
}

// ContextLinkDetailsDTO represents enriched link info with VS Code URI
type ContextLinkDetailsDTO struct {
	ID            string       `json:"id"`
	TaskID        string       `json:"task_id"`
	TaskTitle     string       `json:"task_title"`
	APIContractID string       `json:"api_contract_id"`
	APIContract   *APIContract `json:"api_contract,omitempty"`
	FilePath      string       `json:"file_path"`
	GitBranch     string       `json:"git_branch"`
	VSCodeURI     string       `json:"vscode_uri"`
}

// ContextLinkerRepositoryInterface defines contract for link persistence
type ContextLinkerRepositoryInterface interface {
	CreateLink(ctx context.Context, link *ContextLink) error
	GetLinksByTaskID(ctx context.Context, taskID string) ([]ContextLink, error)
	GetLinkByID(ctx context.Context, id string) (*ContextLink, error)
}
