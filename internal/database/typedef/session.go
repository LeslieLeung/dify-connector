package typedef

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	UserIdentifier string `gorm:"column:user_identifier;unique;primary_key"`
	IsAdmin        bool   `gorm:"column:is_admin"`
	IsBanned       bool   `gorm:"column:is_banned"`
	State          State  `gorm:"column:state;type:longtext;serializer:json"` // JSON encoded state
}

func (s *Session) TableName() string {
	return "session"
}

type State struct {
	CurrentApp int `json:"current_app"`
}
