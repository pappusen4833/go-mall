package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-mall/app/models"
	"go-mall/app/services/dept_service"
	"go-mall/pkg/app"
	"go-mall/pkg/constant"
	"net/http"
)

// 部门api
type DeptController struct {
}

// @Title 获取部门列表
// @Description 获取部门列表
// @Success 200 {object} app.Response
// @router / [get]
// @Tags Admin
func (e *DeptController) GetAll(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	name := c.DefaultQuery("name", "")
	enabled := com.StrTo(c.DefaultQuery("enabled", "-1")).MustInt()
	deptService := dept_service.Dept{Name: name, Enabled: enabled}
	vo := deptService.GetAll()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)
}

// @Title 添加部门
// @Description 添加部门
// @Success 200 {object} app.Response
// @router / [post]
// @Tags Admin
func (e *DeptController) Post(c *gin.Context) {
	var (
		model models.SysDept
		appG  = app.Gin{C: c}
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	deptService := dept_service.Dept{
		M: &model,
	}

	if err := deptService.Insert(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @Title 修改部门
// @Description 修改部门
// @Success 200 {object} app.Response
// @router / [put]
// @Tags Admin
func (e *DeptController) Put(c *gin.Context) {
	var (
		model models.SysDept
		appG  = app.Gin{C: c}
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	deptService := dept_service.Dept{
		M: &model,
	}
	if err := deptService.Save(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}

// @Title 删除部门
// @Description 删除部门
// @Success 200 {object} app.Response
// @router / [delete]
// @Tags Admin
func (e *DeptController) Delete(c *gin.Context) {
	var (
		ids  []int64
		appG = app.Gin{C: c}
	)
	c.BindJSON(&ids)
	deptService := dept_service.Dept{Ids: ids}

	if err := deptService.Del(); err != nil {
		appG.Response(http.StatusInternalServerError, constant.FAIL_ADD_DATA, nil)
		return
	}

	appG.Response(http.StatusOK, constant.SUCCESS, nil)
}
