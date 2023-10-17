package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mall/app/listen"
	"go-mall/app/models"
	"go-mall/packages/base"
	"go-mall/packages/global"
	"go-mall/packages/jwt"
	"go-mall/packages/logging"
	"go-mall/packages/redis"
	"go-mall/packages/wechat"
	"go-mall/routers"
	"log"
	"net/http"
	"time"
)

func init() {
	global.YSHOP_VP = base.Viper()
	global.YSHOP_LOG = base.SetupLogger()
	models.Setup()
	logging.Setup()
	redis.Setup()
	jwt.Setup()
	listen.Setup()
	wechat.InitWechat()
}

// @title gin-shop  API
// @version 1.0
// @description gin-shop商城后台管理系统
// @termsOfService https://gitee.com/guchengwuyue/gin-shop
// @license.name apache2
func main() {
	gin.SetMode(global.YSHOP_CONFIG.Server.RunMode)

	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", global.YSHOP_CONFIG.Server.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}

	global.YSHOP_LOG.Info("[info] start http server listening %s", endPoint)
	log.Printf("[info] start http server listening %s", endPoint)
	fmt.Printf(`
	欢迎使用 go-mall
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	默认前端文件运行地址:http://127.0.0.1:8080
`, endPoint)

	server.ListenAndServe()

}
