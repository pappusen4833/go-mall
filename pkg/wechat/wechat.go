package wechat

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"go-mall/pkg/global"
)

func InitWechat() {
	wc := wechat.NewWechat()
	//这里本地内存保存access_token，也可选择redis，memcache或者自定cache
	redisOpts := &cache.RedisOpts{
		Host:        global.GOMALL_CONFIG.Redis.Host,
		Password:    global.GOMALL_CONFIG.Redis.Password,
		Database:    0,
		MaxActive:   global.GOMALL_CONFIG.Redis.MaxActive,
		MaxIdle:     global.GOMALL_CONFIG.Redis.MaxIdle,
		IdleTimeout: 200,
	}
	redisCache := cache.NewRedis(redisOpts)
	wc.SetCache(redisCache)
	cfg := &offConfig.Config{
		AppID:          global.GOMALL_CONFIG.Wechat.AppID,
		AppSecret:      global.GOMALL_CONFIG.Wechat.AppSecret,
		Token:          global.GOMALL_CONFIG.Wechat.Token,
		EncodingAESKey: global.GOMALL_CONFIG.Wechat.EncodingAESKey,
	}

	officialAccount := wc.GetOfficialAccount(cfg)

	global.GOMALL_OFFICIAL_ACCOUNT = officialAccount
}
