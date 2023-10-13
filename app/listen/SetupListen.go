package listen

import (
	"fmt"
	"go-mall/packages/global"
)

func Setup() {
	var sub PSubscriber
	fmt.Printf(global.YSHOP_CONFIG.Redis.Host)
	conn := PConnect(global.YSHOP_CONFIG.Redis.Host, global.YSHOP_CONFIG.Redis.Password)
	sub.ReceiveKeySpace(conn)
	sub.Psubscribe()
}
