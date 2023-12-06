package app

import (
	"github.com/astaxie/beego/validation"
	"go-mall/pkg/global"
)

// MarkErrors logs error logs
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		global.GOMALL_LOG.Info(err.Key, err.Message)
	}

	return
}
