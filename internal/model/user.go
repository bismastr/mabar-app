package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
	DiscordUID int       `gorm:"unique" json:"discord_uid"`
	Username   string    `json:"username"`
	AvatarURL  string    `json:"avatar_url"`
	Sessions   []Session `gorm:"foreignKey:CreatedBy;references:DiscordUID"`
}
