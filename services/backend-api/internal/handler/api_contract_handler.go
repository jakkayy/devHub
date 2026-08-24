package handler

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jakkayy/devHub/services/backend-api/internal/usecase"
)

type APIContractHandler struct {
	usecase *usecase.APIContractUsecase
}

func NewAPIContractHandler(uc *usecase.APIContractUsecase) *APIContractHandler {
	return &APIContractHandler{usecase: uc}
}

// GetContracts handles GET /api/v1/contracts
func (h *APIContractHandler) GetContracts(c *gin.Context) {
	contracts, err := h.usecase.SyncAndGetContracts(c.Request.Context(), nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"count":  len(contracts),
		"data":   contracts,
	})
}

// GetContractByID handles GET /api/v1/contracts/:id
func (h *APIContractHandler) GetContractByID(c *gin.Context) {
	id := c.Param("id")
	contract, err := h.usecase.GetContractByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "API contract not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   contract,
	})
}

// UploadSpec handles POST /api/v1/contracts/upload (Accepts raw OpenAPI Spec JSON)
func (h *APIContractHandler) UploadSpec(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil || len(body) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid openapi spec json body"})
		return
	}

	contracts, err := h.usecase.SyncAndGetContracts(c.Request.Context(), body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "API contract specification uploaded and parsed successfully",
		"count":   len(contracts),
		"data":    contracts,
	})
}
