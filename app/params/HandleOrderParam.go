package params

import (
	"github.com/astaxie/beego/validation"
)

type HandleOrderParam struct {
	Id string `json:"id"`
}

func (p *HandleOrderParam) Valid(v *validation.Validation) {
	if vv := v.Required(p.Id, "gomall-warning"); !vv.Ok {
		vv.Message("参数有误")
		return
	}
}
