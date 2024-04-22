package middleware

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/leslieleung/dify-connector/internal/log"
)

func WithLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := log.WithRequestID(requestid.Get(c))
		c.Request = c.Request.WithContext(log.ContextWithLogger(c.Request.Context(), logger))
		c.Next()
	}
}
