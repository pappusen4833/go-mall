package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-mall/app/models"
	"go-mall/app/services/dict_service"
	"go-mall/pkg/app"
	"go-mall/pkg/constant"
	"go-mall/pkg/http/response"
	"go-mall/pkg/util"
	"net/http"
)

// 字典api
type DictController struct {
}

// @Title 获取字典列表
// @Description 获取字典列表
// @Success 200 {object} app.Response
// @router /admin/dict [get]
// @Tags Admin
func (e *DictController) GetAll(c *gin.Context) {
	enabled := com.StrTo(c.DefaultQuery("enabled", "-1")).MustInt()
	name := c.DefaultQuery("blurry", "")
	dictService := dict_service.Dict{
		Enabled:  enabled,
		Name:     name,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
	}
	vo := dictService.GetAll()
	response.OkWithData(vo, c)
}

// @Title 添加字典
// @Description 添加字典
// @Success 200 {object} app.Response
// @router /admin/dict [post]
// @Tags Admin
func (e *DictController) Post(c *gin.Context) {
	var (
		model models.SysDict
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	dictService := dict_service.Dict{
		M: &model,
	}

	if err := dictService.Insert(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 修改字典
// @Description 修改字典
// @Success 200 {object} app.Response
// @router /admin/dict [put]
// @Tags Admin
func (e *DictController) Put(c *gin.Context) {
	var (
		model models.SysDict
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	dictService := dict_service.Dict{
		M: &model,
	}

	if err := dictService.Save(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 删除字典
// @Description 删除字典
// @Success 200 {object} app.Response
// @router /admin/dict/:id [delete]
// @Tags Admin
func (e *DictController) Delete(c *gin.Context) {
	var (
		ids []int64
	)
	id := com.StrTo(c.Param("id")).MustInt64()
	ids = append(ids, id)

	dictService := dict_service.Dict{Ids: ids}
	if err := dictService.Del(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}
