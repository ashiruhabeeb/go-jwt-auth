package server

import (
	"go-jwt-auth/internal/db"
	"go-jwt-auth/pkg/config"
	"go-jwt-auth/pkg/logger"
)

func Run() {
	logger := logger.NewSlogHandler()

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Error("[ERROR] loading environment variables failure")
	}
	logger.Info("[INFO] environment variables loaded")

	db, err := db.NewDatabase(cfg.DbDriver, cfg.DbURL)
	if err != nil {
		logger.Error("[ERROR]: %v", err)
	}

	logger.Info("[INFO]: database connection successful")

	db.CloseDB()
}
