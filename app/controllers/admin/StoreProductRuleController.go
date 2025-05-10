package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-mall/app/services/product_rule_service"
	dto2 "go-mall/app/services/product_service/dto"
	"go-mall/pkg/app"
	"go-mall/pkg/constant"
	"go-mall/pkg/http/response"
	"go-mall/pkg/util"
	"net/http"
)

// 商品规格sku api
type StoreProductRuleController struct {
}

// @Title 商品规格sku列表
// @Description 商品规格sku列表
// @Success 200 {object} app.Response
// @router /shop/rule [get]
// @Tags Admin
func (e *StoreProductRuleController) GetAll(c *gin.Context) {
	enabled := com.StrTo(c.DefaultQuery("enabled", "-1")).MustInt()
	name := c.DefaultQuery("blurry", "")
	ruleService := product_rule_service.Rule{
		Enabled:  enabled,
		Name:     name,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
	}
	vo := ruleService.GetAll()
	response.OkWithData(vo, c)
}

// @Title 商品规格sku添加
// @Description 商品规格sku添加
// @Success 200 {object} app.Response
// @router /shop/rule/save/:id [post]
// @Tags Admin
func (e *StoreProductRuleController) Post(c *gin.Context) {
	var (
		dto dto2.ProductRule
	)
	id := com.StrTo(c.Param("id")).MustInt64()
	httpCode, errCode := app.BindAndValid(c, &dto)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	ruleService := product_rule_service.Rule{
		Dto: dto,
		Id:  id,
	}

	if err := ruleService.AddOrSave(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)

}

// @Title 商品规格sku删除
// @Description 商品规格sku删除
// @Success 200 {object} app.Response
// @router /shop/rule [delete]
// @Tags Admin
func (e *StoreProductRuleController) Delete(c *gin.Context) {
	var (
		ids []int64
	)
	c.BindJSON(&ids)
	ruleService := product_rule_service.Rule{Ids: ids}

	if err := ruleService.Del(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}
