package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mall/app/events"
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
	"os"
	"time"
)

func init() {
	global.VP = base.Viper()
	global.LOG = base.SetupLogger()
	models.Setup()
	logging.Setup()
	redis.Setup()
	jwt.Setup()
	listen.Setup()
	wechat.InitWechat()
	events.Setup()
}

// @title GO-MALL API
// @version 1.0
// @description go-mall 商城后台系统
// @license.name apache2
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	gin.SetMode(global.CONFIG.Server.RunMode)

	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", global.CONFIG.Server.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}

	global.LOG.Info("[info] start http server listening %s", server.Addr)
	log.Printf("[info] start http server listening %s", endPoint)

	// Print banner from file
	if banner, err := os.ReadFile("banner.txt"); err == nil {
		fmt.Println(string(banner))
	} else {
		// Fallback banner if file not found
		fmt.Println("GO-MALL")
	}

	fmt.Printf(`
欢迎使用 go-mall
默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
默认前端文件运行地址:http://127.0.0.1:8080
`, endPoint)

	server.ListenAndServe()

}
