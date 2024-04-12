package typedef

import "gorm.io/gorm"

type DifyApp struct {
	gorm.Model
	Name    string `gorm:"column:name"`
	Type    int    `gorm:"column:type"`
	BaseURL string `gorm:"column:base_url"`
	APIKey  string `gorm:"column:api_key"`
	Enabled bool   `gorm:"column:enabled"`
}

func (a *DifyApp) TableName() string {
	return "dify_app"
}
