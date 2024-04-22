package middleware

import (
	"context"
	"github.com/gin-contrib/graceful"
	"github.com/leslieleung/dify-connector/internal/database"
	healthcheck "github.com/tavsec/gin-healthcheck"
	"github.com/tavsec/gin-healthcheck/checks"
	healthcheckcfg "github.com/tavsec/gin-healthcheck/config"
)

func RegisterHealthCheck(r *graceful.Graceful, ctx context.Context) {
	c := []checks.Check{
		checks.NewContextCheck(ctx, "signals"),
		databaseCheck{},
	}
	err := healthcheck.New(r.Engine, healthcheckcfg.DefaultConfig(), c)
	if err != nil {
		return
	}
}

var _ checks.Check = (*databaseCheck)(nil)

type databaseCheck struct{}

func (d databaseCheck) Pass() bool {
	conn := database.GetDB(context.TODO())
	sqlDB, err := conn.DB()
	return err == nil && sqlDB.Ping() == nil
}

func (d databaseCheck) Name() string {
	return "database"
}
