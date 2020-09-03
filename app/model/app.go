package model

import (
	"time"
)

type App struct {
	Id        uint      `gorm:"primary_key"`
	AppId     string    `json:"app_id"`
	AppSecret string    `json:"app_secret"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a App) TableName() string {
	return "app"
}
