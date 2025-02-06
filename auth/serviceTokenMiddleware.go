package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(expectedToken string, errorFactoryFn func(errorMessage string) any) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorFactoryFn("Authorization header отсутствует"))
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorFactoryFn("Неверный формат Authorization header"))
			return
		}

		token := parts[1]
		if token != expectedToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorFactoryFn("Неверный токен"))
			return
		}

		c.Next()
	}
}
