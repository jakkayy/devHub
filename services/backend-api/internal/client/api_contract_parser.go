package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jakkayy/devHub/services/backend-api/internal/domain"
)

type APIContractParser struct{}

func NewAPIContractParser() *APIContractParser {
	return &APIContractParser{}
}

// ParseSpec parses raw OpenAPI 3.0 JSON specification into APIContractDTO slice
func (p *APIContractParser) ParseSpec(ctx context.Context, rawSpec []byte) ([]domain.APIContractDTO, error) {
	if len(rawSpec) == 0 {
		return p.getMockContracts(), nil
	}

	var openAPI struct {
		OpenAPI string `json:"openapi"`
		Swagger string `json:"swagger"`
		Info    struct {
			Version string `json:"version"`
		} `json:"info"`
		Paths map[string]map[string]struct {
			Summary     string   `json:"summary"`
			Description string   `json:"description"`
			Tags        []string `json:"tags"`
		} `json:"paths"`
	}

	if err := json.Unmarshal(rawSpec, &openAPI); err != nil {
		return nil, fmt.Errorf("failed to unmarshal openapi spec json: %w", err)
	}

	specVersion := openAPI.OpenAPI
	if specVersion == "" {
		specVersion = openAPI.Swagger
	}

	var contracts []domain.APIContractDTO
	for pathStr, methods := range openAPI.Paths {
		for methodStr, op := range methods {
			upperMethod := strings.ToUpper(methodStr)
			tag := "General"
			if len(op.Tags) > 0 {
				tag = op.Tags[0]
			}

			contractID := fmt.Sprintf("%s-%s", upperMethod, strings.ReplaceAll(strings.Trim(pathStr, "/"), "/", "-"))
			contracts = append(contracts, domain.APIContractDTO{
				ID:          contractID,
				Method:      upperMethod,
				Path:        pathStr,
				Summary:     op.Summary,
				Description: op.Description,
				SpecVersion: specVersion,
				Tag:         tag,
			})
		}
	}

	return contracts, nil
}

func (p *APIContractParser) getMockContracts() []domain.APIContractDTO {
	return []domain.APIContractDTO{
		{ID: "GET-api-v1-tasks", Method: "GET", Path: "/api/v1/tasks", Summary: "Fetch all tasks synced from Google Sheets and GitHub", SpecVersion: "3.0.0", Tag: "Tasks"},
		{ID: "GET-api-v1-github-pulls", Method: "GET", Path: "/api/v1/github/pulls", Summary: "Fetch active Pull Requests from GitHub", SpecVersion: "3.0.0", Tag: "GitHub"},
		{ID: "GET-api-v1-contracts", Method: "GET", Path: "/api/v1/contracts", Summary: "Fetch interactive OpenAPI specifications", SpecVersion: "3.0.0", Tag: "API Spec"},
	}
}
