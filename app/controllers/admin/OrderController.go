package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/ttlv/kdniao"
	"github.com/ttlv/kdniao/sdk"
	"github.com/unknwon/com"
	"go-mall/app/models"
	"go-mall/app/services/article_service"
	"go-mall/app/services/order_service"
	"go-mall/app/services/order_service/dto"
	"go-mall/pkg/app"
	"go-mall/pkg/constant"
	"go-mall/pkg/global"
	"go-mall/pkg/util"
	"net/http"
)

// order api
type OrderController struct {
}

// @Title 订单列表
// @Description 订单列表
// @Success 200 {object} app.Response
// @router /shop/order [get]
// @Tags Admin
func (e *OrderController) GetAll(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	enabled := com.StrTo(c.DefaultQuery("enabled", "-1")).MustInt()
	name := c.DefaultQuery("blurry", "")
	orderService := order_service.Order{
		Enabled:  enabled,
		Name:     name,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
		IntType:  com.StrTo(c.Query("orderStatus")).MustInt(),
	}
	vo := orderService.GetAll()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}

// @Title 创建订单
// @Description 创建订单
// @Success 200 {object} app.Response
// @router /shop/order [post]
// @Tags Admin
func (e *OrderController) Post(c *gin.Context) {
	var (
		model models.WechatArticle
		appG  = app.Gin{C: c}
	)

	paramErr := app.BindAndValidate(c, &model)
	if paramErr != nil {
		appG.Response(http.StatusBadRequest, paramErr.Error(), nil)
		return
	}

	articleService := article_service.Article{
		M: &model,
	}

	if err := articleService.Insert(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)

}

// @Title 订单修改
// @Description 订单修改
// @Success 200 {object} app.Response
// @router /shop/order [put]
// @Tags Admin
func (e *OrderController) Put(c *gin.Context) {
	var (
		model models.StoreOrder
		appG  = app.Gin{C: c}
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	orderService := order_service.Order{
		M: &model,
	}

	if err := orderService.Save(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @Title 订单发货
// @Description 订单发货
// @Success 200 {object} app.Response
// @router /shop/deliver [put]
// @Tags Admin
func (e *OrderController) Deliver(c *gin.Context) {
	var (
		model models.StoreOrder
		appG  = app.Gin{C: c}
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	orderService := order_service.Order{
		M: &model,
	}

	if err := orderService.Deliver(); err != nil {
		global.LOG.Error(err)
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @Title 订单快递查询
// @Description 订单快递查询
// @Success 200 {object} app.Response
// @router /shop/order/express [post]
// @Tags Admin
func (e *OrderController) DeliverQuery(c *gin.Context) {
	var (
		model dto.Express
		appG  = app.Gin{C: c}
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	config := kdniao.NewKdniaoConfig(global.CONFIG.Express.EBusinessId, global.CONFIG.Express.AppKey)
	logger := kdniao.NewKdniaoLogger()

	expressQuerySdk := sdk.NewExpressQuery(config, logger)
	req := expressQuerySdk.GetRequest(model.LogisticCode)
	resp, err := expressQuerySdk.GetResponse(req)

	if err != nil {
		global.LOG.Error(err)
	}
	//
	if resp.Success == false {
		appG.Response(http.StatusInternalServerError, resp.Reason, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, resp.Traces)
}

// @Title 订单删除
// @Description 订单删除
// @Success 200 {object} app.Response
// @router /shop/order/:id [delete]
// @Tags Admin
func (e *OrderController) Delete(c *gin.Context) {
	var (
		ids  []int64
		appG = app.Gin{C: c}
	)
	id := com.StrTo(c.Param("id")).MustInt64()
	ids = append(ids, id)
	orderService := order_service.Order{Ids: ids}

	if err := orderService.Del(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}
