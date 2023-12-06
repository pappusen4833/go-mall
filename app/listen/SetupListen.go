package listen

import (
	"fmt"
	"go-mall/pkg/global"
)

func Setup() {
	var sub PSubscriber
	fmt.Printf(global.GOMALL_CONFIG.Redis.Host)
	conn := PConnect(global.GOMALL_CONFIG.Redis.Host, global.GOMALL_CONFIG.Redis.Password)
	sub.ReceiveKeySpace(conn)
	sub.Psubscribe()
}
