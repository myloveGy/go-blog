package model

import (
	"fmt"
	"time"

	"blog/global"
	"blog/pkg/setting"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDBEngine(databaseSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	db, err := gorm.Open(databaseSetting.DBType, fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSetting.UserName,
		databaseSetting.Password,
		databaseSetting.Host,
		databaseSetting.DBName,
		databaseSetting.Charset,
		databaseSetting.ParseTime,
	))

	if err != nil {
		return nil, err
	}

	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}

	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:update_at", updateAtForCreateCallback)
	db.Callback().Update().Replace("gorm:update_at", updateAtForUpdateCallback)
	db.DB().SetMaxIdleConns(databaseSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(databaseSetting.MaxOpenConns)

	return db, nil
}

func updateAtForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		if createdAtField, ok := scope.FieldByName("created_at"); ok && createdAtField.IsBlank {
			_ = createdAtField.Set(currentTime)
		}

		if updateAtField, ok := scope.FieldByName("updated_at"); ok && updateAtField.IsBlank {
			_ = updateAtField.Set(currentTime)
		}
	}
}

func updateAtForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("updated_at", time.Now().Format("2006-01-02 15:04:05"))
	}
}
