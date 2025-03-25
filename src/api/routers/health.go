package routers

import (
	"golang-web-api/api/handlers"
	"golang-web-api/api/middlewares"

	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthResponse()
	r.GET("/", middlewares.Limitter() ,handler.Health)
	r.POST("/", handler.Create)
	r.GET("/:id", handler.GetById)
}
