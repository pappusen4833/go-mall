package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mall/app/listen"
	"go-mall/app/models"
	"go-mall/pkg/base"
	"go-mall/pkg/global"
	"go-mall/pkg/jwt"
	"go-mall/pkg/logging"
	"go-mall/pkg/redis"
	"go-mall/pkg/wechat"
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

// @title GO-MALL API
// @version 1.0
// @description go-mall 商城后台系统
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

	global.YSHOP_LOG.Info("[info] start http server listening %s", server.Addr)
	log.Printf("[info] start http server listening %s", endPoint)
	fmt.Println(`
 _____ _____     _____ _____ __    __    
|   __|     |___|     |  _  |  |  |  |   
|  |  |  |  |___| | | |     |  |__|  |__ 
|_____|_____|   |_|_|_|__|__|_____|_____|
`)
	fmt.Printf(`

欢迎使用 go-mall
默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
默认前端文件运行地址:http://127.0.0.1:8080
`, endPoint)

	server.ListenAndServe()

}
