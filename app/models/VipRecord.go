package models

import "go-mall/pkg/global"

type VipRecord struct {
	Uid         int64  `gorm:"column:uid;default:NULL" json:"userId"`
	OrderNo     string `gorm:"column:order_no;default:NULL" json:"orderNo"`
	ExpiredTime int64  `gorm:"column:expired_time;default:NULL" json:"expiredTime"`
	Vip         int8   `gorm:"column:vip;default:0" json:"vip"`
	BaseModel
}

func (m *VipRecord) TableName() string {
	return "vip_record"
}

func AddVipRecord(m *VipRecord) error {
	var err error
	if err = db.Create(m).Error; err != nil {
		return err
	}
	return err
}

func FindVipRecordByUserId(uid int64) VipRecord {
	var record VipRecord
	err := db.Where("uid = ?", uid).First(&record).Error
	if err != nil {
		global.LOG.Info(err)
	}
	return record
}

func GetAllVipRecord(pageNUm int, pageSize int, maps interface{}) (int64, []VipRecord) {
	var (
		total int64
		list  []VipRecord
	)
	db.Model(&VipRecord{}).Where(maps).Count(&total)
	db.Where(maps).Offset(pageNUm).Limit(pageSize).Find(&list)
	return total, list
}

func UpdateByVipRecord(m *VipRecord) error {
	var err error
	err = db.Save(m).Error
	if err != nil {
		return err
	}
	return err
}
func DelByVipRecord(ids []int64) error {
	var err error
	err = db.Where("id in (?)", ids).Delete(&VipRecord{}).Error
	if err != nil {
		return err
	}
	return err
}
