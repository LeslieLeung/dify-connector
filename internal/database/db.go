package database

import (
	"context"
	"github.com/leslieleung/dify-connector/internal/database/typedef"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(dsn string) {
	db = New(dsn)
}

func New(dsn string) *gorm.DB {
	database, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}))
	if err != nil {
		panic(err)
	}
	// Ping to verify the connection
	conn, err := database.DB()
	if err != nil {
		panic(err)
	}
	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	// auto migrate
	err = database.AutoMigrate(&typedef.Channel{}, &typedef.DifyApp{}, &typedef.Session{})
	if err != nil {
		panic(err)
	}

	return database
}

func GetDB(ctx context.Context) *gorm.DB {
	return db.WithContext(ctx)
}
