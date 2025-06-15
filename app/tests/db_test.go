package tests

import (
	"go-mall/app/models"
	"go-mall/pkg/base"
	"go-mall/pkg/global"
	"testing"
)

func init() {
	testing.Init()
	global.VP = base.Viper("../../config.yaml")
	//global.LOG = base.SetupLogger()
	models.Setup()
	//global.DB, _ = gorm.Open(mysql.Open("chatbot:chatbot1212121@tcp(127.0.0.1:3306)/chatbot?charset=utf8&parseTime=True&loc=Local"))
}

func Test(t *testing.T) {
}
