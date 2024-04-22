package log

import (
	"context"
	"go.uber.org/zap"
)

type ctxLogger struct{}

func WithRequestID(requestID string) *zap.Logger {
	return Default().With(zap.String("request_id", requestID))
}

func ContextWithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, logger)
}

func FromContext(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value(ctxLogger{}).(*zap.Logger)
	if !ok {
		return Default()
	}
	return logger
}
