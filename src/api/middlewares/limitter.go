package middlewares

import (
	"net/http"
	"golang-web-api/api/helper"
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)

func Limitter() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(1, nil)
	return func(c *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, helper.GenerateBaseResponseWithError(nil, false , -100 , err))
			return
		}else{
			c.Next()
		}
	}
}
