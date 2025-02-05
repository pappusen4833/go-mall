package events

import (
	"github.com/asaskevich/EventBus"
	"go-mall/pkg/global"
)

func Setup() {
	global.EVENT_BUS = EventBus.New()
	_ = global.EVENT_BUS.Subscribe("user:create", OnUserCreate)
}
