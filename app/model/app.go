package model

import (
	"time"

	"github.com/jinzhu/gorm"
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

func (a *App) Get(db *gorm.DB) (*App, error) {
	err := db.Where("app_id = ? AND app_secret = ?", a.AppId, a.AppSecret).First(a).Error
	return a, err
}
