package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jakkayy/devHub/services/backend-api/internal/domain"
)

type GitHubClient struct {
	httpClient *http.Client
	patToken   string
}

func NewGitHubClient(patToken string) *GitHubClient {
	return &GitHubClient{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		patToken:   patToken,
	}
}

// FetchPullRequests fetches open Pull Requests from a GitHub repository
func (c *GitHubClient) FetchPullRequests(ctx context.Context, owner string, repo string) ([]domain.GitHubPullRequestDTO, error) {
	// Fallback mock data if token or repo is not configured
	if c.patToken == "" || owner == "" || repo == "" {
		return c.getMockPullRequests(), nil
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls?state=open", owner, repo)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create github pr request: %w", err)
	}

	c.setHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch github pull requests: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github api returned status: %d", resp.StatusCode)
	}

	var rawPulls []struct {
		ID        int64     `json:"id"`
		Number    int       `json:"number"`
		Title     string    `json:"title"`
		State     string    `json:"state"`
		HTMLURL   string    `json:"html_url"`
		Draft     bool      `json:"draft"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		User      struct {
			Login string `json:"login"`
		} `json:"user"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&rawPulls); err != nil {
		return nil, fmt.Errorf("failed to decode github pulls json: %w", err)
	}

	var prs []domain.GitHubPullRequestDTO
	for _, item := range rawPulls {
		prs = append(prs, domain.GitHubPullRequestDTO{
			ID:        item.ID,
			Number:    item.Number,
			Title:     item.Title,
			State:     item.State,
			Author:    item.User.Login,
			HTMLURL:   item.HTMLURL,
			Draft:     item.Draft,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	return prs, nil
}

// FetchIssues fetches open Issues from a GitHub repository
func (c *GitHubClient) FetchIssues(ctx context.Context, owner string, repo string) ([]domain.GitHubIssueDTO, error) {
	if c.patToken == "" || owner == "" || repo == "" {
		return c.getMockIssues(), nil
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues?state=open", owner, repo)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create github issues request: %w", err)
	}

	c.setHeaders(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch github issues: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github api returned status: %d", resp.StatusCode)
	}

	var rawIssues []struct {
		ID        int64     `json:"id"`
		Number    int       `json:"number"`
		Title     string    `json:"title"`
		State     string    `json:"state"`
		HTMLURL   string    `json:"html_url"`
		CreatedAt time.Time `json:"created_at"`
		Assignee  *struct {
			Login string `json:"login"`
		} `json:"assignee"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&rawIssues); err != nil {
		return nil, fmt.Errorf("failed to decode github issues json: %w", err)
	}

	var issues []domain.GitHubIssueDTO
	for _, item := range rawIssues {
		assigneeName := "Unassigned"
		if item.Assignee != nil {
			assigneeName = item.Assignee.Login
		}
		issues = append(issues, domain.GitHubIssueDTO{
			ID:        item.ID,
			Number:    item.Number,
			Title:     item.Title,
			State:     item.State,
			Assignee:  assigneeName,
			HTMLURL:   item.HTMLURL,
			CreatedAt: item.CreatedAt,
		})
	}

	return issues, nil
}

func (c *GitHubClient) setHeaders(req *http.Request) {
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("User-Agent", "devHub-Platform")
	if c.patToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.patToken)
	}
}

func (c *GitHubClient) getMockPullRequests() []domain.GitHubPullRequestDTO {
	return []domain.GitHubPullRequestDTO{
		{ID: 1001, Number: 1, Title: "feat(backend): setup postgres database & clean architecture", State: "open", Author: "jakkayy", HTMLURL: "https://github.com/jakkayy/devHub/pull/1", Draft: false, CreatedAt: time.Now().Add(-2 * time.Hour), UpdatedAt: time.Now()},
		{ID: 1002, Number: 2, Title: "ci: setup github actions pipeline", State: "open", Author: "jakkayy", HTMLURL: "https://github.com/jakkayy/devHub/pull/2", Draft: false, CreatedAt: time.Now().Add(-1 * time.Hour), UpdatedAt: time.Now()},
	}
}

func (c *GitHubClient) getMockIssues() []domain.GitHubIssueDTO {
	return []domain.GitHubIssueDTO{
		{ID: 2001, Number: 10, Title: "Add Redis Caching layer for external tools", State: "open", Assignee: "jakkayy", HTMLURL: "https://github.com/jakkayy/devHub/issues/10", CreatedAt: time.Now().Add(-5 * time.Hour)},
	}
}
