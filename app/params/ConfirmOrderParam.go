package params

import (
	"github.com/astaxie/beego/validation"
)

type ConfirmOrderParam struct {
	CartId string `json:"cartId"`
}

func (p *ConfirmOrderParam) Valid(v *validation.Validation) {
	if vv := v.Required(p.CartId, "gomall-warning"); !vv.Ok {
		vv.Message("提交购买的商品")
		return
	}
}
