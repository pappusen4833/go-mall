package vip_record_service

import (
	"errors"
	"go-mall/app/models"
	"go-mall/app/models/vo"
	"go-mall/pkg/global"
	"gorm.io/gorm"
	"time"
)

type VipRecord struct {
	User *models.User

	Id      int64
	Name    string
	Enabled int

	PageNum  int
	PageSize int

	M *models.VipRecord

	Ids []int64
}

// CreateNew 开通会员记录
func (e *VipRecord) UpsertVipRecord(period int, vip int8) {
	var record models.VipRecord
	err := global.DB.Model(&models.VipRecord{}).Where("uid = ?", e.User.Id).First(&record).Error
	startTime := time.Now().Unix()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// new insert
		var newRecord models.VipRecord
		newRecord.Uid = e.User.Id
		newRecord.Vip = vip
		newRecord.ExpiredTime = startTime + int64(period)
		err = global.DB.Create(&newRecord).Error
		global.LOG.Info(err)

	} else {
		// old, update current vip time
		if time.Now().Unix() < record.ExpiredTime {
			startTime = record.ExpiredTime
			record.ExpiredTime = record.ExpiredTime + int64(period)
			record.Vip = vip
			err = global.DB.Save(&record).Error
			global.LOG.Info(err)
		}
	}
}

func (d *VipRecord) GetAll() vo.ResultList {
	maps := make(map[string]interface{})
	if d.Enabled >= 0 {
		maps["enabled"] = d.Enabled
	}
	if d.Name != "" {
		maps["name"] = d.Name
	}

	total, list := models.GetAllVipRecord(d.PageNum, d.PageSize, maps)
	return vo.ResultList{Content: list, TotalElements: total}
}

func (d *VipRecord) Insert() error {
	return models.AddVipRecord(d.M)
}

func (d *VipRecord) Save() error {
	return models.UpdateByVipRecord(d.M)
}

func (d *VipRecord) Del() error {
	return models.DelByVipRecord(d.Ids)
}
