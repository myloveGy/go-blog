package global

import (
	"blog/pkg/logger"
	"blog/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
	JwtSetting      *setting.JwtSettingS
)
