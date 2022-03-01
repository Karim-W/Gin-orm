package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AuthenticationMiddleware(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
