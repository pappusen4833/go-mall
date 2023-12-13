package global

import (
	"github.com/silenceper/wechat/v2/officialaccount"
	"github.com/spf13/viper"
	"go-mall/conf"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB               *gorm.DB
	VP               *viper.Viper
	LOG              *zap.SugaredLogger
	CONFIG           conf.Config
	OFFICIAL_ACCOUNT *officialaccount.OfficialAccount
)
