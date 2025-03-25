package handlers

import (
	"errors"
	"fmt"
	"golang-web-api/api/helper"
	"golang-web-api/data/cache"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthResponse struct {	
}

func NewHealthResponse() *HealthResponse {
	return &HealthResponse{}
}

// @Summary Health Check
// @Description Check if the server is running
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.BaseHttpResponse
// @Router /v1/health [get]
func (h *HealthResponse) Health(c *gin.Context) {
	redisClient := cache.GetRedisClient()
	
	c.JSON(200, helper.GenerateBaseResponse(gin.H{
		"redis": redisClient.Ping().String(),
	}, true, 200))
}

func (h *HealthResponse) Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"test": "create new data",
	})
}

// @Summary Get Data By Id
// @Description Get data by id
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.BaseHttpResponse
// @Failure 400 {object} map[string]any
// @Param id path string true "id"
// @Router /v1/health/{id} [get]
// @Security AuthBearer
func (h *HealthResponse) GetById(c *gin.Context) {
	id := c.Param("id")
	
	if id == "5" {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithError(nil, false, 400, errors.New("id is required")))
		return
	}
	
	c.JSON(200, gin.H{
		"test": fmt.Sprintf("get data by id %s", id),
	})
}