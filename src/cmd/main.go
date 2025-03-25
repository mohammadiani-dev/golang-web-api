package main

import (
	"golang-web-api/api"
	"golang-web-api/config"
	"golang-web-api/data/cache"
	"golang-web-api/data/db"
	"golang-web-api/pkg/logging"
)

//@securityDefinitions.apikey AuthBearer
//@in header
//@name Authorization
func main() {
	cfg := config.GetConfig()

	logger := logging.NewLogger(&cfg.Logger)

	err := cache.InitRedis(&cfg.Redis)
	if err != nil {
		logger.Fatal(logging.Redis , logging.StartUp , err.Error(), nil)
	}
	defer cache.CloseRedis()

	err = db.InitDb(&cfg.Postgres)
	if err != nil {
		logger.Fatal(logging.Postgres , logging.StartUp , err.Error(), nil)
	}
	defer db.CloseDb()

	api.InitServer(cfg)
}