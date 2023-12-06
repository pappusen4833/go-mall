package params

import "github.com/astaxie/beego/validation"

type RegParam struct {
	Account  string `json:"account"`
	Captcha  string `json:"captcha"`
	Password string `json:"password"`
	Spread   string `json:"spread"`
}

func (p *RegParam) Valid(v *validation.Validation) {
	if vv := v.Phone(p.Account, "gomall-warning"); !vv.Ok {
		vv.Message("手机格式不对")
		return
	}
	if vv := v.Required(p.Captcha, "gomall-warning"); !vv.Ok {
		vv.Message("验证码必填")
		return
	}
	if vv := v.Required(p.Password, "gomall-warning"); !vv.Ok {
		vv.Message("密码必填")
		return
	}

}
