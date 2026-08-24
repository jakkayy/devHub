package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jakkayy/devHub/services/backend-api/internal/domain"
	"github.com/jakkayy/devHub/services/backend-api/internal/usecase"
)

type ContextLinkerHandler struct {
	usecase *usecase.ContextLinkerUsecase
}

func NewContextLinkerHandler(uc *usecase.ContextLinkerUsecase) *ContextLinkerHandler {
	return &ContextLinkerHandler{usecase: uc}
}

// CreateLink handles POST /api/v1/links
func (h *ContextLinkerHandler) CreateLink(c *gin.Context) {
	var dto domain.CreateLinkDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	linkDetails, err := h.usecase.CreateContextLink(c.Request.Context(), dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Context link created successfully",
		"data":    linkDetails,
	})
}

// GetLinksByTaskID handles GET /api/v1/links/task/:task_id
func (h *ContextLinkerHandler) GetLinksByTaskID(c *gin.Context) {
	taskID := c.Param("task_id")
	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task_id path parameter is required"})
		return
	}

	links, err := h.usecase.GetLinksForTask(c.Request.Context(), taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"count":  len(links),
		"data":   links,
	})
}
