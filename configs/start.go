package configs

import (
	"log"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"

	"blog/global"
	"blog/pkg/logger"
	"blog/pkg/model"
	"blog/pkg/setting"
)

func init() {
	if err := setupSetting(); err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	log.Println("init.setupSetting success")

	if err := setupDBEngine(); err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	log.Println("init.setupDBEngine success")

	if err := setupLogger(); err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

func setupSetting() error {
	configSetting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	if err = configSetting.ReadSection("Server", &global.ServerSetting); err != nil {
		return err
	}

	if err = configSetting.ReadSection("App", &global.AppSetting); err != nil {
		return err
	}

	if err = configSetting.ReadSection("Database", &global.DatabaseSetting); err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	return err
}

func setupLogger() error {
	global.Logger = logger.NewLogger(
		&lumberjack.Logger{
			Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
			MaxSize:   600,
			MaxAge:    10,
			LocalTime: true,
		},
		"",
		log.LstdFlags,
	).WithCaller(2)

	return nil
}
