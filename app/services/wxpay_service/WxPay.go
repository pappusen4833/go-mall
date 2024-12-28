package wxpay_service

import (
	"encoding/json"
	"fmt"
	WXPay "github.com/wleven/wxpay"
	"github.com/wleven/wxpay/src/V2"
	"github.com/wleven/wxpay/src/entity"
	"go-mall/app/models"
	"go-mall/app/services/vip_record_service"
	"go-mall/pkg/global"
	"time"
)

type WxPay struct {
	User *models.User
}

// WeChat Pay configuration - should be loaded from config file or environment variables
var (
	wxpayapi WXPay.WXPayApi
)

func init() {
	// Load WeChat Pay configuration from global config
	config := entity.PayConfig{
		AppID:     global.CONFIG.Wechat.AppID,
		MchID:     global.CONFIG.WxPay.MchID,
		PayNotify: global.CONFIG.WxPay.NotifyURL,
		Secret:    global.CONFIG.WxPay.APIKey,
	}
	wxpayapi = WXPay.Init(config)
}

func (e *WxPay) Query(outTradeNo string) {
	query, err := wxpayapi.V2.OrderQuery(V2.OrderQuery{OutTradeNo: outTradeNo})
	if err != nil {
		return
	}
	global.LOG.Info(query)
}

func (e *WxPay) Notify(data string) bool {
	notify := V2.NotifyFormat(data)
	returnCode := notify["return_code"]
	resultCode := notify["result_code"]
	outTradeNo := notify["out_trade_no"]
	if returnCode == "SUCCESS" && resultCode == "SUCCESS" {
		// update order status
		payOrder := models.VipPayOrder{}
		err := global.DB.Model(&models.VipPayOrder{}).Where("order_no = ?", outTradeNo).First(&payOrder).Error
		if err != nil {
			return false
		}
		if payOrder.Status == 1 {
			global.LOG.Info("already resolved")
			return false
		}
		payOrder.Status = 1
		global.DB.Save(&payOrder)
		// give user vip time
		product := models.VipProduct{}
		_ = json.Unmarshal([]byte(payOrder.ProductInfo), &product)
		global.LOG.Info(product)
		user := models.User{}
		user.Id = payOrder.Uid
		record := vip_record_service.VipRecord{User: &user}
		record.UpsertVipRecord(product.Period, 1)
	}
	global.LOG.Info(outTradeNo)
	return true
}

func (e *WxPay) Pay(productId int64, method string) *map[string]interface{} {
	memberId := e.User.Id
	global.LOG.Info("user,%s,try to buy product,%s\n", memberId, productId)
	product := models.GetVipProduct(productId)
	outTradeNo := GenerateOrderNo()
	payAmount := product.Price
	orderName := product.Name
	productInfo, _ := json.Marshal(product)
	var (
		payOrder = models.VipPayOrder{
			OrderNo:     outTradeNo,
			OrderName:   orderName,
			TotalAmount: payAmount,
			Uid:         memberId,
			Status:      0,
			ProductInfo: string(productInfo),
		}
	)
	global.DB.Save(&payOrder)
	if method == "native" {
		data, err := wxpayapi.V2.UnifiedOrder(V2.UnifiedOrder{
			OutTradeNo:     outTradeNo,
			TotalFee:       int(payAmount),
			SpbillCreateIP: "127.0.0.1",
			NotifyURL:      global.CONFIG.WxPay.NotifyURL,
			Attach:         "pay",
			TradeType:      "NATIVE",
			Body:           orderName,
		})
		if err != nil {
			global.LOG.Info(err)
		}
		global.LOG.Info(data)
		global.LOG.Info(data["code_url"])

		// Generate QR code URL using configured service
		qrServiceURL := global.CONFIG.WxPay.QRServiceURL
		if qrServiceURL == "" {
			qrServiceURL = "https://api.qrserver.com/v1/create-qr-code/?size=300x300&data=" // Default free service
		}
		url := qrServiceURL + data["code_url"]
		return &map[string]interface{}{
			"outTradeNo": outTradeNo,
			"payAmount":  payAmount,
			"url":        url,
			"code_url":   data["code_url"],
			"returnUrl":  global.CONFIG.WxPay.ReturnURL,
		}
	} else {
		// JSAPI payment mode (for WeChat in-app payment)
		// This is legacy code that needs to be refactored
		//scheme := "https://"
		//wxPay := NewWxJsApiPayService(MCH_ID, APP_ID, APP_KEY, API_KEY)
		//uri := r.URL.RequestURI()
		//frontendURL := global.CONFIG.App.FrontendBaseURL
		//redirectUrl := frontendURL + uri
		//openId := wxuser.OpenId
		//if openId == "" {
		//	response := Result{Code: 500, Message: "获取openid失败"}
		//	json.NewEncoder(w).Encode(response)
		//	return
		//}
		//notifyURL := global.CONFIG.WxPay.NotifyURL
		//jsApiParameters := wxPay.CreateJsBizPackage(openId, payAmount, outTradeNo, orderName, notifyURL, payTime)
		//jsonData, _ := json.Marshal(jsApiParameters)
		//data := map[string]interface{}{
		//	"outTradeNo":      outTradeNo,
		//	"jsApiParameters": string(jsonData),
		//	"payAmount":       payAmountStr,
		//}
		//response := Result{Code: 200, Data: data}
		//json.NewEncoder(w).Encode(response)
	}
	return nil
}

func (w *WxPay) QueryOrder(outTradeNo string) bool {
	result, err := wxpayapi.V2.OrderQuery(V2.OrderQuery{
		OutTradeNo: outTradeNo,
	})
	if err != nil {
		return false
	}
	if result["return_code"] != "SUCCESS" {
		return false
	}
	if result["trade_state"] != "SUCCESS" {
		return false
	}
	global.LOG.Info(result)
	return true
}

func GenerateOrderNo() string {
	// implementation
	now := time.Now()
	sec := now.Unix()
	usec := now.UnixNano() % 0x100000
	return fmt.Sprintf("wx%08x%05x", sec, usec)
}
