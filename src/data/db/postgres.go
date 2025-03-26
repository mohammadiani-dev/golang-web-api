package db

import (
	"fmt"
	"golang-web-api/config"
	"golang-web-api/pkg/logging"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

var logger = logging.NewLogger(&config.GetConfig().Logger)


func InitDb(cfg *config.Postgres) error {
	var err error
	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	dbClient, err = gorm.Open(postgres.Open(cnn), &gorm.Config{})
	if err != nil {
		return err
	}
	
	sqlDB, _ := dbClient.DB()
	err = sqlDB.Ping()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime * time.Minute)

	logger.Info(logging.Postgres , logging.StartUp , "Connected to database", nil)

	return nil
}

func GetDbClient() *gorm.DB {
	return dbClient
}

func CloseDb() error {
	sqlDB, _ := dbClient.DB()
	return sqlDB.Close()
}

