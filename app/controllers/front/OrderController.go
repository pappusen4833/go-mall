package front

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay/wechat"
	"github.com/unknwon/com"
	"go-mall/app/params"
	"go-mall/app/services/order_service"
	orderDto "go-mall/app/services/order_service/dto"
	"go-mall/app/services/pay_service"
	"go-mall/pkg/app"
	"go-mall/pkg/global"
	"go-mall/pkg/http/response"
	"go-mall/pkg/jwt"
	"go-mall/pkg/util"
	"net/http"
)

// Order api
type OrderController struct {
}

// @Title 订单确认
// @Description 订单确认
// @Success 200 {object} app.Response
// @router /api/v1/order/confirm [post]
// @Tags Front API
func (e *OrderController) Confirm(c *gin.Context) {
	var (
		param params.ConfirmOrderParam
	)
	paramErr := app.BindAndValidate(c, &param)
	if paramErr != nil {
		response.Error(http.StatusBadRequest, 9999, paramErr.Error(), nil, c)
		return
	}

	uid, _ := jwt.GetAppUserId(c)
	user, _ := jwt.GetAppDetailUser(c)
	orderService := order_service.Order{
		CartId: param.CartId,
		Uid:    uid,
		User:   user,
	}
	vo, err := orderService.ConfirmOrder()
	if err != nil {
		response.Error(http.StatusInternalServerError, 9999, err.Error(), nil, c)
		return
	}
	response.OkWithData(vo, c)

}

// @Title 订单计算
// @Description 订单计算
// @Success 200 {object} app.Response
// @router /api/v1/order/computed/:key [post]
// @Tags Front API
func (e *OrderController) Compute(c *gin.Context) {
	var (
		param params.ComputeOrderParam
	)
	paramErr := app.BindAndValidate(c, &param)
	if paramErr != nil {
		response.Error(http.StatusBadRequest, 9999, paramErr.Error(), nil, c)
		return
	}

	uid, _ := jwt.GetAppUserId(c)
	//user,_:= jwt.GetAppDetailUser(c)
	orderService := order_service.Order{
		Uid:          uid,
		ComputeParam: &param,
		Key:          c.Param("key"),
	}
	checkMap, err := orderService.Check()
	if err != nil {
		response.Error(http.StatusInternalServerError, 9999, err.Error(), nil, c)
		return
	}
	if checkMap != nil {
		response.OkWithDetailed(checkMap, checkMap["msg"].(string), c)
		return
	}
	vo, err := orderService.ComputeOrder()
	if err != nil {
		response.Error(http.StatusInternalServerError, 9999, err.Error(), nil, c)
		return
	}
	response.OkWithData(gin.H{
		"result": vo,
		"status": "NONE",
	}, c)
}

// @Title 订单创建
// @Description 订单创建
// @Success 200 {object} app.Response
// @router /api/v1/order/create/:key [post]
// @Tags Front API
func (e *OrderController) Create(c *gin.Context) {
	var (
		param params.OrderParam
	)
	paramErr := app.BindAndValidate(c, &param)
	if paramErr != nil {
		response.Error(http.StatusBadRequest, 9999, paramErr.Error(), nil, c)
		return
	}

	uid, _ := jwt.GetAppUserId(c)
	//user,_:= jwt.GetAppDetailUser(c)
	key := c.Param("key")
	orderService := order_service.Order{
		Uid:        uid,
		OrderParam: &param,
		Key:        key,
	}
	order, err := orderService.CreateOrder()
	if err != nil {
		response.Error(http.StatusInternalServerError, 9999, err.Error(), nil, c)
		return
	}

	global.LOG.Info(order)
	orderExtendDto := &orderDto.OrderExtend{
		Key:     key,
		OrderId: order.OrderId,
	}
	returnMap := gin.H{
		"status":     "SUCCESS",
		"result":     orderExtendDto,
		"createTune": order.CreateTime,
	}

	response.OkWithData(returnMap, c)

}

// @Title 订单支付
// @Description 订单支付
// @Success 200 {object} app.Response
// @router /api/v1/order/pay [post]
// @Tags Front API
func (e *OrderController) Pay(c *gin.Context) {
	var (
		param params.PayParam
	)
	paramErr := app.BindAndValidate(c, &param)
	if paramErr != nil {
		response.Error(http.StatusBadRequest, 9999, paramErr.Error(), nil, c)
		return
	}

	uid, _ := jwt.GetAppUserId(c)
	orderExtendDto := &orderDto.OrderExtend{
		//Key: key,
		OrderId: param.Uni,
	}
	returnMap := gin.H{
		"status": "SUCCESS",
		"result": orderExtendDto,
	}

	newMap, err := pay_service.GoPay(returnMap, param.Uni, param.PayType, param.From, uid, orderExtendDto)
	if err != nil {
		response.Error(http.StatusInternalServerError, 9999, err.Error(), nil, c)
		return
	}
	response.OkWithData(newMap, c)

}

// @Title 订单移步支付
// @Description 订单移步支付
// @Success 200 {object} app.Response
// @router /api/v1/order/notify [get]
// @Tags Front API
func (e *OrderController) NotifyPay(c *gin.Context) {
	notifyReq, err := wechat.ParseNotifyToBodyMap(c.Request)
	//支付成功后处理
	if err != nil {
		global.LOG.Error(err)
	}

	global.LOG.Info(notifyReq)

}

// @Title 订单详情
// @Description 订单详情
// @Success 200 {object} app.Response
// @router /api/v1/order/detail/:key [get]
// @Tags Front API
func (e *OrderController) OrderDetail(c *gin.Context) {
	uid, _ := jwt.GetAppUserId(c)
	//user,_:= jwt.GetAppDetailUser(c)
	key := c.Param("key")
	orderService := order_service.Order{
		Uid: uid,
		//OrderParam: &param,
		OrderId: key,
	}
	order, _ := orderService.GetOrderInfo()

	newOrder := order_service.HandleOrder(order)

	response.OkWithData(newOrder, c)

}

// @Title 获取列表数据
// @Description 获取列表数据
// @Success 200 {object} app.Response
// @router /api/v1/order [get]
// @Tags Front API
func (e *OrderController) GetList(c *gin.Context) {
	uid, _ := jwt.GetAppUserId(c)
	orderService := order_service.Order{
		IntType:  com.StrTo(c.Query("type")).MustInt(),
		PageNum:  util.GetFrontPage(c),
		PageSize: util.GetFrontLimit(c),
		Uid:      uid,
	}

	vo, total, page := orderService.GetList()
	response.PageResult(0, vo, "ok", total, page, c)
}

// @Title 订单收货
// @Description 订单收货
// @Success 200 {object} app.Response
// @router /api/v1/order/take [post]
// @Tags Front API
func (e *OrderController) TakeOrder(c *gin.Context) {
	var (
		param params.DoOrderParam
	)
	paramErr := app.BindAndValidate(c, &param)
	if paramErr != nil {
		response.Error(http.StatusBadRequest, 9999, paramErr.Error(), nil, c)
		return
	}

	uid, _ := jwt.GetAppUserId(c)
	orderService := order_service.Order{
		OrderId: param.Uni,
		Uid:     uid,
	}

	if err := orderService.TakeOrder(); err != nil {
		response.Error(http.StatusInternalServerError, 9999, err.Error(), nil, c)
		return
	}
	response.OkWithData("success", c)

}

// @Title 订单评价
// @Description 订单评价
// @Success 200 {object} app.Response
// @router /api/v1/order/comments/:key [post]
// @Tags Front API
func (e *OrderController) OrderComment(c *gin.Context) {
	var (
		param []params.ProductReplyParam
	)
	c.ShouldBindJSON(&param)

	uid, _ := jwt.GetAppUserId(c)
	orderService := order_service.Order{
		OrderId:    c.Param("key"),
		Uid:        uid,
		ReplyParam: param,
	}

	if err := orderService.OrderComment(); err != nil {
		response.Error(http.StatusInternalServerError, 9999, err.Error(), nil, c)
		return
	}
	response.OkWithData("ok", c)

}

// @Title 未支付订单取消
// @Description 未支付订单取消
// @Success 200 {object} app.Response
// @router /api/v1/order/cancel [post]
// @Tags Front API
func (e *OrderController) CancelOrder(c *gin.Context) {
	var (
		param params.HandleOrderParam
	)
	paramErr := app.BindAndValidate(c, &param)
	if paramErr != nil {
		response.Error(http.StatusBadRequest, 9999, paramErr.Error(), nil, c)
		return
	}

	uid, _ := jwt.GetAppUserId(c)
	orderService := order_service.Order{
		OrderId: param.Id,
		Uid:     uid,
	}

	if err := orderService.CancelOrder(); err != nil {
		response.Error(http.StatusInternalServerError, 9999, err.Error(), nil, c)
		return
	}
	response.OkWithData("success", c)

}
