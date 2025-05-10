package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-mall/app/models"
	dto2 "go-mall/app/services/menu_service/dto"
	"go-mall/app/services/role_service"
	"go-mall/pkg/app"
	"go-mall/pkg/constant"
	"go-mall/pkg/http/response"
	"go-mall/pkg/logging"
	"go-mall/pkg/util"
	"net/http"
)

// 角色 API
type RoleController struct {
}

// @Title 获取单个角色
// @Description 获取单个角色
// @Param    id        path     int    true        "角色ID"
// @Success 200 {object} app.Response
// @router /admin/roles/:id [get]
// @Tags Admin
func (e *RoleController) GetOne(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt64()
	roleService := role_service.Role{
		Id: id,
	}
	vo := roleService.GetOneRole()
	response.OkWithData(vo, c)
}

// @Title 角色列表
// @Description 角色列表
// @Success 200 {object} app.Response
// @router /admin/roles [get]
// @Tags Admin
func (e *RoleController) GetAll(c *gin.Context) {
	blurry := c.DefaultQuery("blurry", "")
	roleService := role_service.Role{
		Name:     blurry,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
	}
	vo := roleService.GetAll()
	response.OkWithData(vo, c)
}

// @Title 角色添加
// @Description 角色添加
// @Success 200 {object} app.Response
// @router /admin/roles [post]
// @Tags Admin
func (e *RoleController) Post(c *gin.Context) {
	var (
		model models.SysRole
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	roleService := role_service.Role{
		M: &model,
	}

	if err := roleService.Insert(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @router /admin/roles [put]
// @Tags Admin
func (e *RoleController) Put(c *gin.Context) {
	var (
		model models.SysRole
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	roleService := role_service.Role{
		M: &model,
	}

	if err := roleService.Save(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 角色删除
// @Description 角色删除
// @Success 200 {object} app.Response
// @router /admin/roles [delete]
// @Tags Admin
func (e *RoleController) Delete(c *gin.Context) {
	var (
		ids []int64
	)
	c.BindJSON(&ids)
	roleService := role_service.Role{Ids: ids}

	if err := roleService.Del(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 角色菜单更新
// @Description 角色菜单更新
// @Success 200 {object} app.Response
// @router /admin/roles/menu [put]
// @Tags Admin
func (e *RoleController) Menu(c *gin.Context) {
	var (
		model dto2.RoleMenu
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	logging.Info(model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}

	roleService := role_service.Role{Dto: model}
	if err := roleService.BatchRoleMenuAdd(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)

}
