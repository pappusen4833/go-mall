package events

import (
	"fmt"
	"go-mall/app/models"
	"go-mall/app/services/vip_record_service"
	"go-mall/pkg/global"
)

func OnUserCreate(userId string) {
	fmt.Println("On User Create Event...")
	var user models.User
	err := global.DB.
		Model(&models.User{}).
		Where("username = ?", userId).First(&user).Error
	if err != nil {
		return
	}
	record := vip_record_service.VipRecord{User: &user}
	//record.CreateNew(30, 30, 0, 0)
	record.UpsertVipRecord(1*86400, 0)
	global.LOG.Info("create member record finished")
}
