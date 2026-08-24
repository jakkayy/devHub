package domain

import "time"

// Task represents a task synced from Google Sheets or GitHub Issues
type Task struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Assignee    string    `json:"assignee"`
	Status      string    `json:"status"` // e.g. To-Do, In Progress, Done
	Source      string    `json:"source"` // sheets, github
	ExternalURL string    `json:"external_url"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// APIContract represents an API Specification endpoint
type APIContract struct {
	ID          string `json:"id" gorm:"primaryKey"`
	Method      string `json:"method"` // GET, POST, PUT, DELETE
	Path        string `json:"path"`
	Summary     string `json:"summary"`
	SpecVersion string `json:"spec_version"`
}

// ContextLink connects Task, Code File, and API Contract together
type ContextLink struct {
	ID            string    `json:"id" gorm:"primaryKey"`
	TaskID        string    `json:"task_id" index:"idx_task_link"`
	APIContractID string    `json:"api_contract_id" index:"idx_api_link"`
	FilePath      string    `json:"file_path"`
	GitBranch     string    `json:"git_branch"`
	CreatedAt     time.Time `json:"created_at"`
}
