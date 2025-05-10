package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-mall/app/models"
	"go-mall/app/services/material_group_service"
	"go-mall/pkg/app"
	"go-mall/pkg/constant"
	"go-mall/pkg/http/response"
	"go-mall/pkg/jwt"
	"net/http"
)

// 素材分组api
type MaterialGroupController struct {
}

// @Title 素材分组列表
// @Description 素材分组列表
// @Success 200 {object} app.Response
// @router /admin/materialgroup [get]
// @Tags Admin
func (e *MaterialGroupController) GetAll(c *gin.Context) {
	name := c.DefaultQuery("blurry", "")
	materialGroupService := material_group_service.MaterialGroup{
		Name: name,
	}
	vo := materialGroupService.GetAll()
	response.OkWithData(vo, c)
}

// @Title素材分组添加
// @Description素材分组添加
// @Success 200 {object} app.Response
// @router /admin/materialgroup [post]
// @Tags Admin
func (e *MaterialGroupController) Post(c *gin.Context) {
	var (
		model models.SysMaterialGroup
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	uid, _ := jwt.GetAdminUserId(c)
	model.CreateId = uid
	materialGroupService := material_group_service.MaterialGroup{
		M: &model,
	}

	if err := materialGroupService.Insert(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)

}

// @Title 素材分组修改
// @Description 素材分组修改
// @Success 200 {object} app.Response
// @router /admin/materialgroup [put]
// @Tags Admin
func (e *MaterialGroupController) Put(c *gin.Context) {
	var (
		model models.SysMaterialGroup
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	uid, _ := jwt.GetAdminUserId(c)
	model.CreateId = uid
	materialGroupService := material_group_service.MaterialGroup{
		M: &model,
	}

	if err := materialGroupService.Save(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 素材分组删除
// @Description 素材分组删除
// @Success 200 {object} app.Response
// @router /admin/materialgroup/:id [delete]
// @Tags Admin
func (e *MaterialGroupController) Delete(c *gin.Context) {
	var (
		ids []int64
	)
	id := com.StrTo(c.Param("id")).MustInt64()
	ids = append(ids, id)
	materialGroupService := material_group_service.MaterialGroup{Ids: ids}

	if err := materialGroupService.Del(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}
