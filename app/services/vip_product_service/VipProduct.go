package vip_product_service

import (
	"go-mall/app/models"
	"go-mall/app/models/vo"
)

type VipProduct struct {
	Id      int64
	Name    string
	Enabled int

	PageNum  int
	PageSize int

	M *models.VipProduct

	Ids []int64
}

func (d *VipProduct) GetAll() vo.ResultList {
	maps := make(map[string]interface{})
	if d.Enabled >= 0 {
		maps["enabled"] = d.Enabled
	}
	if d.Name != "" {
		maps["name"] = d.Name
	}

	total, list := models.GetAllVipProduct(d.PageNum, d.PageSize, maps)
	return vo.ResultList{Content: list, TotalElements: total}
}

func (d *VipProduct) Insert() error {
	return models.AddVipProduct(d.M)
}

func (d *VipProduct) Save() error {
	return models.UpdateByVipProduct(d.M)
}

func (d *VipProduct) Del() error {
	return models.DelByVipProduct(d.Ids)
}
