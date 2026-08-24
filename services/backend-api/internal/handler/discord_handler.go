package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jakkayy/devHub/services/backend-api/internal/domain"
	"github.com/jakkayy/devHub/services/backend-api/internal/usecase"
)

type DiscordHandler struct {
	usecase *usecase.DiscordUsecase
}

func NewDiscordHandler(uc *usecase.DiscordUsecase) *DiscordHandler {
	return &DiscordHandler{usecase: uc}
}

// SendNotification handles POST /api/v1/discord/notify
func (h *DiscordHandler) SendNotification(c *gin.Context) {
	var dto domain.NotificationEventDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.SendEventNotification(c.Request.Context(), dto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Discord notification dispatched successfully",
		"data":    dto,
	})
}
