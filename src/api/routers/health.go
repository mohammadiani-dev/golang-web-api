package routers

import (
	"golang-web-api/api/handlers"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthResponse()
	r.GET("/", handler.Health)
	r.POST("/", handler.Create)
	r.GET("/:id", handler.GetById)
}
