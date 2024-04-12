package database

import (
	"context"
	"github.com/leslieleung/dify-connector/internal/database/typedef"
	"gorm.io/gorm"
)

func CreateDifyApp(ctx context.Context, app *typedef.DifyApp) error {
	return GetDB(ctx).Create(app).Error
}

func GetDifyApp(ctx context.Context, id int) (*typedef.DifyApp, error) {
	app := &typedef.DifyApp{}
	err := GetDB(ctx).Where("id = ?", id).First(app).Error
	return app, err
}

func GetDifyApps(ctx context.Context) ([]*typedef.DifyApp, error) {
	var apps []*typedef.DifyApp
	err := GetDB(ctx).Find(&apps).Error
	return apps, err
}

func GetEnabledApps(ctx context.Context) ([]*typedef.DifyApp, error) {
	var apps []*typedef.DifyApp
	err := GetDB(ctx).Where("enabled = ?", true).Find(&apps).Error
	return apps, err
}

func ToggleApp(ctx context.Context, id int) error {
	return GetDB(ctx).Model(&typedef.DifyApp{}).
		Where("id = ?", id).
		Update("enabled", gorm.Expr("NOT enabled")).
		Error
}

func RemoveApp(ctx context.Context, id int) error {
	return GetDB(ctx).Delete(&typedef.DifyApp{}, id).Error
}
