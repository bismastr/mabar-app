package model

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	ID      int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name    string    `json:"name"`
	IconURL string    `json:"icon_url"`
	Session []Session `gorm:"foreignKey:GameId"`
}

type Session struct {
	gorm.Model
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"`
	IsFinish     bool      `gorm:"default:false" json:"is_finish"`
	CreatedAt    time.Time `json:"created_at"`
	SessionEnd   time.Time `json:"session_end"`
	SessionStart time.Time `json:"session_start"`
	CreatedBy    int
	GameId       int
	User         User `gorm:"foreignKey:CreatedBy;references:DiscordUID" json:"user"`
	Game         Game `gorm:"foreignKey:GameId;references:ID" json:"game"`
}
