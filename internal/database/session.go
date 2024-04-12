package database

import (
	"context"
	"github.com/leslieleung/dify-connector/internal/database/typedef"
)

func GetSession(ctx context.Context, uid string) (*typedef.Session, error) {
	session := &typedef.Session{}
	err := GetDB(ctx).Where("user_identifier = ?", uid).First(session).Error
	return session, err
}

func SaveSession(ctx context.Context, session *typedef.Session) error {
	return GetDB(ctx).Save(session).Error
}
