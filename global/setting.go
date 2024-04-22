package global

import (
	"github.com/wekeeroad/GoWeb/pkg/logger"
	"github.com/wekeeroad/GoWeb/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
	JWTSetting      *setting.JWTSetting
)
