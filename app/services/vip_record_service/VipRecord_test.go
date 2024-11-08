package vip_record_service

import (
	"errors"
	"go-mall/app/models"
	"go-mall/pkg/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_VipRecord(t *testing.T) {
	db, err := gorm.Open(mysql.Open("chatbot:chatbot1212121@tcp(127.0.0.1:3306)/chatbot?charset=utf8&parseTime=True&loc=Local"))
	var record models.VipRecord
	err = db.Model(&models.VipRecord{}).Where("uid = ?", 39).First(&record).Error
	global.LOG.Info(errors.Is(err, gorm.ErrRecordNotFound))
}
