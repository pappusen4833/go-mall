package global

import (
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/spf13/viper"
	"go-mall/conf"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GOMALL_DB               *gorm.DB
	GOMALL_VP               *viper.Viper
	GOMALL_LOG              *zap.SugaredLogger
	GOMALL_CONFIG           conf.Config
	GOMALL_OFFICIAL_ACCOUNT *officialaccount.OfficialAccount
)
