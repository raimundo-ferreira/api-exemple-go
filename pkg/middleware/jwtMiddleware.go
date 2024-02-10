package middleware

import (
	"api-exemple/app/data/response"
	"api-exemple/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.Msg{Message: "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
