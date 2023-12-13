package listen

import (
	"fmt"
	"go-mall/pkg/global"
)

func Setup() {
	var sub PSubscriber
	fmt.Printf(global.CONFIG.Redis.Host)
	conn := PConnect(global.CONFIG.Redis.Host, global.CONFIG.Redis.Password)
	sub.ReceiveKeySpace(conn)
	sub.Psubscribe()
}
