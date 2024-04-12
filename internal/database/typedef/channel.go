package typedef

import "gorm.io/gorm"

type Channel struct {
	gorm.Model
	Name       string `gorm:"column:name"`
	Type       int    `gorm:"column:type"`
	Credential string `gorm:"column:credential"` // channel specific credentials, encoded in JSON
	Enabled    bool   `gorm:"column:enabled"`
}

func (c *Channel) TableName() string {
	return "channel"
}
