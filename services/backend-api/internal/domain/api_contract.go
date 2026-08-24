package domain

import "context"

// APIContractDTO represents detailed API Specification info for Interactive Explorer
type APIContractDTO struct {
	ID          string `json:"id"`
	Method      string `json:"method"` // GET, POST, PUT, DELETE
	Path        string `json:"path"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	SpecVersion string `json:"spec_version"`
	Tag         string `json:"tag"`
}

// APIContractParserInterface defines contract for parsing OpenAPI specs
type APIContractParserInterface interface {
	ParseSpec(ctx context.Context, rawSpec []byte) ([]APIContractDTO, error)
}

// APIContractRepositoryInterface defines contract for database operations on API Specs
type APIContractRepositoryInterface interface {
	UpsertContracts(ctx context.Context, contracts []APIContract) error
	GetAllContracts(ctx context.Context) ([]APIContract, error)
	GetContractByID(ctx context.Context, id string) (*APIContract, error)
}
