package api

import (
	"fmt"
	"golang-web-api/api/middlewares"
	"golang-web-api/api/routers"
	"golang-web-api/config"
	"golang-web-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config){
	r := gin.New()

	//middlewares
	r.Use(middlewares.DefaultStructuredLogger(&cfg.Logger))
	r.Use(gin.Logger() , gin.Recovery())

	InitRouter(r)
	RegisterSwagger(r , &cfg.Server)
	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func InitRouter(r *gin.Engine){
	api := r.Group("/api")
	
	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		routers.Health(health)
	}
}

func RegisterSwagger(r *gin.Engine , cfg *config.Server){
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Host = "localhost:" + cfg.Port
	docs.SwaggerInfo.Title = "Car Sale API"
	docs.SwaggerInfo.Description = "API for Car Sale"
	docs.SwaggerInfo.Version = "1.0"
		
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

