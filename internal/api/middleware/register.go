package middleware

import (
	"bytes"
	"github.com/gin-contrib/graceful"
	"github.com/gin-contrib/requestid"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/leslieleung/dify-connector/internal/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"time"
)

func RegisterMiddlewares(r *graceful.Graceful) {
	r.Use(requestid.New())
	r.Use(WithLogger())
	logger := log.Default()
	r.Use(ginzap.GinzapWithConfig(logger, &ginzap.Config{
		TimeFormat: time.RFC3339Nano,
		UTC:        true,
		SkipPaths:  []string{},
		Context: func(c *gin.Context) []zapcore.Field {
			fields := []zapcore.Field{}
			fields = append(fields, zap.String("request_id", requestid.Get(c)))

			// log request body
			var body []byte
			var buf bytes.Buffer
			tee := io.TeeReader(c.Request.Body, &buf)
			body, _ = io.ReadAll(tee)
			c.Request.Body = io.NopCloser(&buf)
			fields = append(fields, zap.String("body", string(body)))
			return fields
		},
	}))
}
