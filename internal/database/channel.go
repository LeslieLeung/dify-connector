package database

import (
	"context"
	"github.com/leslieleung/dify-connector/internal/database/typedef"
)

func CreateChannel(ctx context.Context, channel *typedef.Channel) error {
	return GetDB(ctx).Create(channel).Error
}

func GetChannel(ctx context.Context, id int) (*typedef.Channel, error) {
	channel := &typedef.Channel{}
	err := GetDB(ctx).Where("id = ?", id).First(channel).Error
	return channel, err
}

func GetChannels(ctx context.Context) ([]*typedef.Channel, error) {
	var channels []*typedef.Channel
	err := GetDB(ctx).Find(&channels).Error
	return channels, err
}

func GetEnabledChannels(ctx context.Context) ([]*typedef.Channel, error) {
	var channels []*typedef.Channel
	err := GetDB(ctx).Where("enabled = ?", true).Find(&channels).Error
	return channels, err
}

func ToggleChannel(ctx context.Context, id int, enabled bool) error {
	return GetDB(ctx).Model(&typedef.Channel{}).Where("id = ?", id).Update("enabled", enabled).Error
}

func SaveChannel(ctx context.Context, channel *typedef.Channel) error {
	return GetDB(ctx).Save(channel).Error
}
