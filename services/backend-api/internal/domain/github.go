package domain

import (
	"context"
	"time"
)

// GitHubPullRequestDTO represents a Pull Request from GitHub API
type GitHubPullRequestDTO struct {
	ID        int64     `json:"id"`
	Number    int       `json:"number"`
	Title     string    `json:"title"`
	State     string    `json:"state"`
	Author    string    `json:"author"`
	HTMLURL   string    `json:"html_url"`
	Draft     bool      `json:"draft"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GitHubIssueDTO represents an Issue from GitHub API
type GitHubIssueDTO struct {
	ID        int64     `json:"id"`
	Number    int       `json:"number"`
	Title     string    `json:"title"`
	State     string    `json:"state"`
	Assignee  string    `json:"assignee"`
	HTMLURL   string    `json:"html_url"`
	CreatedAt time.Time `json:"created_at"`
}

// GitHubClientInterface defines external contract for fetching GitHub data
type GitHubClientInterface interface {
	FetchPullRequests(ctx context.Context, owner string, repo string) ([]GitHubPullRequestDTO, error)
	FetchIssues(ctx context.Context, owner string, repo string) ([]GitHubIssueDTO, error)
}
