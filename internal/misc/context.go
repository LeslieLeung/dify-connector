package misc

import (
	"context"
	"github.com/gin-gonic/gin"
)

func ToContext(c *gin.Context) context.Context {
	return c.Request.Context()
}
