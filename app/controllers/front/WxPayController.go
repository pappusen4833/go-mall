package front

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go-mall/app/services/wxpay_service"
	"go-mall/pkg/global"
	"go-mall/pkg/http/response"
	"go-mall/pkg/jwt"
	"strconv"
)

type WxPayController struct {
}

func (e *WxPayController) Pay(c *gin.Context) {
	global.LOG.Info("start pay")
	user, _ := jwt.GetAppDetailUser(c)
	productId, _ := strconv.ParseInt(c.Query("productId"), 10, 64)
	method := c.Query("method")
	pay := wxpay_service.WxPay{User: user}
	response.OkWithData(pay.Pay(productId, method), c)
}

func (e *WxPayController) Query(c *gin.Context) {
	outTradeNo := c.Query("outTradeNo")
	user, _ := jwt.GetAppDetailUser(c)
	pay := wxpay_service.WxPay{User: user}
	res := pay.QueryOrder(outTradeNo)

	if res {
		response.OkWithData(nil, c)
	} else {
		response.FailWithCodeMessage(1, "", c)
	}

}

func (e *WxPayController) Notify(c *gin.Context) {
	global.LOG.Info(">>> wxpay notify")

	user, _ := jwt.GetAppDetailUser(c)
	pay := wxpay_service.WxPay{User: user}
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(c.Request.Body)
	data := buf.String()
	global.LOG.Info(data)
	pay.Notify(data)
	response.OkWithData(nil, c)
}
