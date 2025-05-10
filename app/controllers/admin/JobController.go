package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-mall/app/models"
	"go-mall/app/services/job_service"
	"go-mall/pkg/app"
	"go-mall/pkg/constant"
	"go-mall/pkg/http/response"
	"go-mall/pkg/util"
	"net/http"
)

// 岗位api
type JobController struct {
}

// @Title 岗位列表
// @Description 岗位列表
// @Success 200 {object} app.Response
// @router /admin/job [get]
// @Tags Admin
func (e *JobController) GetAll(c *gin.Context) {
	enabled := com.StrTo(c.DefaultQuery("enabled", "-1")).MustInt()
	name := c.DefaultQuery("blurry", "")
	jobService := job_service.Job{
		Enabled:  enabled,
		Name:     name,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
	}
	vo := jobService.GetAll()
	response.OkWithData(vo, c)
}

// @Title 岗位添加
// @Description 岗位添加
// @Success 200 {object} app.Response
// @router /admin/job [post]
// @Tags Admin
func (e *JobController) Post(c *gin.Context) {
	var (
		model models.SysJob
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	jobService := job_service.Job{
		M: &model,
	}

	if err := jobService.Insert(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)

}

// @Title 岗位修改
// @Description 岗位修改
// @Success 200 {object} app.Response
// @router /admin/job [put]
// @Tags Admin
func (e *JobController) Put(c *gin.Context) {
	var (
		model models.SysJob
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	jobService := job_service.Job{
		M: &model,
	}

	if err := jobService.Save(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 岗位删除
// @Description 岗位删除
// @Success 200 {object} app.Response
// @router /admin/job [delete]
// @Tags Admin
func (e *JobController) Delete(c *gin.Context) {
	var (
		ids []int64
	)
	c.BindJSON(&ids)
	jobService := job_service.Job{Ids: ids}

	if err := jobService.Del(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}
