package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jakkayy/devHub/services/backend-api/internal/domain"
)

type SheetsClient struct {
	httpClient *http.Client
}

func NewSheetsClient() *SheetsClient {
	return &SheetsClient{
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

// FetchTasks fetches data from a public/API-key Google Sheet or Mock JSON fallback
func (c *SheetsClient) FetchTasks(ctx context.Context, spreadsheetID string, sheetRange string) ([]domain.SheetTaskDTO, error) {
	// If spreadsheetID is empty, return initial mock tasks for development
	if spreadsheetID == "" {
		return c.getMockSheetTasks(), nil
	}

	url := fmt.Sprintf("https://sheets.googleapis.com/v4/spreadsheets/%s/values/%s", spreadsheetID, sheetRange)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create sheets request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch sheets data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("sheets api returned status: %d", resp.StatusCode)
	}

	var sheetResp struct {
		Values [][]string `json:"values"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&sheetResp); err != nil {
		return nil, fmt.Errorf("failed to decode sheets json: %w", err)
	}

	var tasks []domain.SheetTaskDTO
	for idx, row := range sheetResp.Values {
		if idx == 0 || len(row) < 4 { // Skip header or incomplete rows
			continue
		}
		tasks = append(tasks, domain.SheetTaskDTO{
			RowIndex: idx,
			TaskID:   row[0],
			Title:    row[1],
			Assignee: row[2],
			Status:   row[3],
		})
	}

	return tasks, nil
}

func (c *SheetsClient) getMockSheetTasks() []domain.SheetTaskDTO {
	return []domain.SheetTaskDTO{
		{RowIndex: 1, TaskID: "TASK-101", Title: "Setup Database & Clean Architecture", Assignee: "Developer", Status: "Done"},
		{RowIndex: 2, TaskID: "TASK-102", Title: "Integrate Google Sheets & GitHub APIs", Assignee: "Developer", Status: "In Progress"},
		{RowIndex: 3, TaskID: "TASK-103", Title: "Build Miro Diagram Embed View", Assignee: "Developer", Status: "To Do"},
	}
}
