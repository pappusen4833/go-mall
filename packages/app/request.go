package app

import (
	"github.com/astaxie/beego/validation"
	"go-mall/packages/global"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		global.YSHOP_LOG.Info(err.Key, err.Message)
	}

	return
}
