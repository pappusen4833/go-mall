package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-mall/app/models"
	"go-mall/app/services/dict_detail_service"
	"go-mall/pkg/app"
	"go-mall/pkg/constant"
	"go-mall/pkg/http/response"
	"go-mall/pkg/util"
	"net/http"
)

// 字典详情api
type DictDetailController struct {
}

// @Title 获取字典详情列表
// @Description 获取字典详情列表
// @Success 200 {object} app.Response
// @router /admin/dictDetail [get]
// @Tags Admin
func (e *DictDetailController) GetAll(c *gin.Context) {
	dictName := c.DefaultQuery("dictName", "")
	dictId := com.StrTo(c.DefaultQuery("dictId", "-1")).MustInt64()
	detailService := dict_detail_service.DictDetail{
		DictName: dictName,
		DictId:   dictId,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
	}
	vo := detailService.GetAll()
	response.OkWithData(vo, c)
}

// @Title 添加字典详情
// @Description 添加字典详情
// @Success 200 {object} app.Response
// @router /admin/dictDetail [post]
// @Tags Admin
func (e *DictDetailController) Post(c *gin.Context) {
	var (
		model models.SysDictDetail
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	dictDetailService := dict_detail_service.DictDetail{
		M: &model,
	}

	if err := dictDetailService.Insert(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 修改字典详情
// @Description 修改字典详情
// @Success 200 {object} app.Response
// @router /admin/dictDetail [put]
// @Tags Admin
func (e *DictDetailController) Put(c *gin.Context) {
	var (
		model models.SysDictDetail
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	dictDetailService := dict_detail_service.DictDetail{
		M: &model,
	}

	if err := dictDetailService.Save(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 删除字典详情
// @Description 删除字典详情
// @Success 200 {object} app.Response
// @router /admin/dictDetail/:id [delete]
// @Tags Admin
func (e *DictDetailController) Delete(c *gin.Context) {
	var (
		ids []int64
	)
	id := com.StrTo(c.Param("id")).MustInt64()
	ids = append(ids, id)

	dictDetailService := dict_detail_service.DictDetail{Ids: ids}
	if err := dictDetailService.Del(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}
