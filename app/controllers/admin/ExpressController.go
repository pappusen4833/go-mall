package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-mall/app/models"
	"go-mall/app/services/express_service"
	"go-mall/pkg/app"
	"go-mall/pkg/constant"
	"go-mall/pkg/http/response"
	"go-mall/pkg/util"
	"net/http"
)

// 快递api
type ExpressController struct {
}

// @Title 快递列表
// @Description 快递列表
// @Success 200 {object} app.Response
// @router /shop/express [get]
// @Tags Admin
func (e *ExpressController) GetAll(c *gin.Context) {
	enabled := com.StrTo(c.DefaultQuery("enabled", "-1")).MustInt()
	name := c.DefaultQuery("blurry", "")
	expressService := express_service.Express{
		Enabled:  enabled,
		Name:     name,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
	}
	vo := expressService.GetAll()
	response.OkWithData(vo, c)
}

// @Title 快递添加
// @Description 快递添加
// @Success 200 {object} app.Response
// @router /shop/express [post]
// @Tags Admin
func (e *ExpressController) Post(c *gin.Context) {
	var (
		model models.Express
	)

	paramErr := app.BindAndValidate(c, &model)
	if paramErr != nil {
		response.Error(http.StatusBadRequest, 9999, paramErr.Error(), nil, c)
		return
	}

	expressService := express_service.Express{
		M: &model,
	}

	if err := expressService.Insert(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)

}

// @Title 快递修改
// @Description 快递修改
// @Success 200 {object} app.Response
// @router /shop/express [put]
// @Tags Admin
func (e *ExpressController) Put(c *gin.Context) {
	var (
		model models.Express
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	expressService := express_service.Express{
		M: &model,
	}

	if err := expressService.Save(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 快递删除
// @Description 快递删除
// @Success 200 {object} app.Response
// @router /shop/express/:id [delete]
// @Tags Admin
func (e *ExpressController) Delete(c *gin.Context) {
	var (
		ids []int64
	)
	id := com.StrTo(c.Param("id")).MustInt64()
	ids = append(ids, id)
	expressService := express_service.Express{Ids: ids}

	if err := expressService.Del(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}
