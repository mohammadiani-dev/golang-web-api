package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type HealthResponse struct {	
}

func NewHealthResponse() *HealthResponse {
	return &HealthResponse{}
}

func (h *HealthResponse) Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"test": "new data",
	})
}

func (h *HealthResponse) Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"test": "create new data",
	})
}

func (h *HealthResponse) GetById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"test": fmt.Sprintf("get data by id %s", id),
	})
}


