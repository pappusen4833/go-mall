package models

import "go-mall/pkg/global"

type VipProduct struct {
	Type      string `json:"type"`
	Level     string `json:"level"`
	Name      string `json:"name"`
	Count     int    `json:"count"`
	Period    int    `json:"period"`
	PeriodStr string `json:"periodStr"`
	Price     int64  `json:"price"`
	BaseModel
}

func (VipProduct) TableName() string {
	return "vip_product"
}

func GetVipProduct(productId int64) VipProduct {
	var product VipProduct
	global.DB.First(&product, productId)
	return product
}

// get all
func GetAllVipProduct(pageNUm int, pageSize int, maps interface{}) (int64, []VipProduct) {
	var (
		total int64
		list  []VipProduct
	)
	db.Model(&VipProduct{}).Where(maps).Count(&total)
	db.Where(maps).Offset(pageNUm).Limit(pageSize).Find(&list)
	return total, list
}

// last inserted Id on success.
func AddVipProduct(m *VipProduct) error {
	var err error
	if err = db.Create(m).Error; err != nil {
		return err
	}
	return err
}
func UpdateByVipProduct(m *VipProduct) error {
	var err error
	err = db.Save(m).Error
	if err != nil {
		return err
	}
	return err
}
func DelByVipProduct(ids []int64) error {
	var err error
	err = db.Where("id in (?)", ids).Delete(&VipProduct{}).Error
	if err != nil {
		return err
	}
	return err
}
