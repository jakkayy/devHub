package domain

import "context"

// SheetTaskDTO represents raw task data parsed from Google Sheets row
type SheetTaskDTO struct {
	RowIndex    int    `json:"row_index"`
	TaskID      string `json:"task_id"`
	Title       string `json:"title"`
	Assignee    string `json:"assignee"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	ExternalURL string `json:"external_url"`
}

// SheetsClientInterface defines external contract for fetching Google Sheets data
type SheetsClientInterface interface {
	FetchTasks(ctx context.Context, spreadsheetID string, sheetRange string) ([]SheetTaskDTO, error)
}
