package models

type VipPayOrder struct {
	Uid         int64  `gorm:"column:uid;NOT NULL"`
	OrderNo     string `gorm:"column:order_no;NOT NULL"`
	OrderName   string `gorm:"column:order_name;default:NULL"`
	ProductId   int32  `gorm:"column:product_id;default:NULL"`
	ProductInfo string `gorm:"column:product_info;default:NULL"`
	Status      int32  `gorm:"column:status;default:0;NOT NULL;comment:'0: created, 1: paid'"`
	TotalAmount int64  `json:"totalAmount"`
	BaseModel
}

func (v *VipPayOrder) TableName() string {
	return "vip_pay_order"
}
