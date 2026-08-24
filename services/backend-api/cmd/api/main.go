package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jakkayy/devHub/services/backend-api/internal/client"
	"github.com/jakkayy/devHub/services/backend-api/internal/config"
	"github.com/jakkayy/devHub/services/backend-api/internal/handler"
	"github.com/jakkayy/devHub/services/backend-api/internal/repository"
	"github.com/jakkayy/devHub/services/backend-api/internal/usecase"
)

func main() {
	// 1. Load Configuration
	cfg := config.LoadConfig()

	// 2. Initialize Repositories (PostgreSQL & Redis)
	db, err := repository.NewPostgresDatabase(cfg)
	if err != nil {
		log.Printf("Warning: Database connection skipped/failed: %v", err)
	}

	taskRepo := repository.NewTaskRepository(db.DB)
	contractRepo := repository.NewAPIContractRepository(db.DB)
	linkerRepo := repository.NewContextLinkerRepository(db.DB)
	cacheRepo := repository.NewCacheRepository(cfg.RedisHost, cfg.RedisPassword)

	// 3. Initialize External Clients & Parsers
	sheetsClient := client.NewSheetsClient()
	githubClient := client.NewGitHubClient(cfg.GitHubPAT)
	contractParser := client.NewAPIContractParser()

	// 4. Initialize Business Logic Usecases
	integrationUsecase := usecase.NewIntegrationUsecase(sheetsClient, githubClient, taskRepo, cacheRepo)
	contractUsecase := usecase.NewAPIContractUsecase(contractParser, contractRepo, cacheRepo)
	linkerUsecase := usecase.NewContextLinkerUsecase(linkerRepo, contractRepo, cacheRepo)

	// 5. Initialize Delivery Handlers
	integrationHandler := handler.NewIntegrationHandler(integrationUsecase, cfg)
	contractHandler := handler.NewAPIContractHandler(contractUsecase)
	linkerHandler := handler.NewContextLinkerHandler(linkerUsecase)

	// 6. Setup Gin Router & Routes
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"service":   "devHub-backend-api",
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})

	api := r.Group("/api/v1")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong from devHub Go Backend Service!",
			})
		})

		// Task & GitHub Endpoints
		api.GET("/tasks", integrationHandler.GetTasks)
		api.GET("/github/pulls", integrationHandler.GetGitHubPullRequests)

		// API Contract Specs Endpoints
		api.GET("/contracts", contractHandler.GetContracts)
		api.GET("/contracts/:id", contractHandler.GetContractByID)
		api.POST("/contracts/upload", contractHandler.UploadSpec)

		// Context Linker & Deep Link Endpoints
		api.POST("/links", linkerHandler.CreateLink)
		api.GET("/links/task/:task_id", linkerHandler.GetLinksByTaskID)
	}

	log.Printf("Starting devHub Go Backend Service on :%s...", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
