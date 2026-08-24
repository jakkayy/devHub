package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jakkayy/devHub/services/backend-api/internal/config"
	"github.com/jakkayy/devHub/services/backend-api/internal/repository"
)

func main() {
	cfg := config.LoadConfig()

	db, err := repository.NewPostgresDatabase(cfg)
	if err != nil {
		log.Printf("Warning: Database connection skipped/failed: %v", err)
	} else {
		_ = db
	}

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
	}

	log.Printf("Starting devHub Go Backend Service on :%s...", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
