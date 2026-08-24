package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jakkayy/devHub/services/backend-api/internal/config"
	"github.com/jakkayy/devHub/services/backend-api/internal/usecase"
)

type IntegrationHandler struct {
	usecase *usecase.IntegrationUsecase
	cfg     *config.Config
}

func NewIntegrationHandler(uc *usecase.IntegrationUsecase, cfg *config.Config) *IntegrationHandler {
	return &IntegrationHandler{
		usecase: uc,
		cfg:     cfg,
	}
}

// GetTasks handles GET /api/v1/tasks
func (h *IntegrationHandler) GetTasks(c *gin.Context) {
	spreadsheetID := c.Query("spreadsheet_id")
	if spreadsheetID == "" {
		spreadsheetID = h.cfg.GoogleSheetID
	}

	tasks, err := h.usecase.SyncAndGetTasks(c.Request.Context(), spreadsheetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"count":  len(tasks),
		"data":   tasks,
	})
}

// GetGitHubPullRequests handles GET /api/v1/github/pulls
func (h *IntegrationHandler) GetGitHubPullRequests(c *gin.Context) {
	owner := c.Query("owner")
	repo := c.Query("repo")

	pulls, err := h.usecase.GetGitHubPullRequests(c.Request.Context(), owner, repo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"count":  len(pulls),
		"data":   pulls,
	})
}
