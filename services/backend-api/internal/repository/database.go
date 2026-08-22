package repository

import (
	"fmt"
	"log"

	"github.com/jakkayy/devHub/services/backend-api/internal/config"
	"github.com/jakkayy/devHub/services/backend-api/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	DB *gorm.DB
}

func NewPostgresDatabase(cfg *config.Config) (*Database, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Bangkok",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode,
	)

	gormConfig := &gorm.Config{}
	if cfg.Environment == "development" {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL database")

	// Run AutoMigrations
	if err := AutoMigrate(db); err != nil {
		return nil, fmt.Errorf("failed to run database auto migrations: %w", err)
	}

	return &Database{DB: db}, nil
}

func AutoMigrate(db *gorm.DB) error {
	log.Println("Running Database AutoMigrations for Domain Models...")
	err := db.AutoMigrate(
		&domain.Task{},
		&domain.APIContract{},
		&domain.ContextLink{},
	)
	if err != nil {
		return err
	}
	log.Println("Database AutoMigrations completed successfully!")
	return nil
}
