package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-mall/app/models"
	"go-mall/app/services/cron_job_service"
	"go-mall/pkg/app"
	"go-mall/pkg/constant"
	"go-mall/pkg/http/response"
	"go-mall/pkg/util"
	"net/http"
)

type SysCronJobController struct {
}

// @Title 获取定时任务调度表列表
// @Description 获取定时任务调度表列表
// @Success 200 {object} app.Response
// @router /tools/timing [get]
// @Tags Admin
func (e *SysCronJobController) GetAll(c *gin.Context) {
	enabled := com.StrTo(c.DefaultQuery("enabled", "-1")).MustInt()
	name := c.DefaultQuery("blurry", "")
	service := cron_job_service.SysCronJob{
		Enabled:  enabled,
		Name:     name,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
	}
	vo := service.GetAll()
	response.OkWithData(vo, c)
}

// @Title 添加定时任务调度表
// @Description 添加定时任务调度表
// @Success 200 {object} app.Response
// @router /tools/timing [post]
// @Tags Admin
func (e *SysCronJobController) Post(c *gin.Context) {
	var (
		model models.SysCronJob
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	service := cron_job_service.SysCronJob{
		M: &model,
	}

	if err := service.Insert(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 修改定时任务调度表
// @Description 修改定时任务调度表
// @Success 200 {object} app.Response
// @router /tools/timing [put]
// @Tags Admin
func (e *SysCronJobController) Put(c *gin.Context) {
	var (
		model models.SysCronJob
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	service := cron_job_service.SysCronJob{
		M: &model,
	}

	if err := service.Save(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 删除定时任务调度表
// @Description 删除定时任务调度表
// @Success 200 {object} app.Response
// @router /tools/timing/:id [delete]
// @Tags Admin
func (e *SysCronJobController) Delete(c *gin.Context) {
	var (
		ids []int64
	)

	if strId := c.Param("id"); strId != "" {
		id := com.StrTo(strId).MustInt64()
		ids = append(ids, id)
	} else {
		c.BindJSON(&ids)
	}

	service := cron_job_service.SysCronJob{Ids: ids}
	if err := service.Del(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 执行定时任务调度表
// @Description 执行定时任务调度表
// @Success 200 {object} app.Response
// @router /tools/timing/exec/:id [put]
// @Tags Admin
func (e *SysCronJobController) Exec(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt64()
	service := cron_job_service.SysCronJob{
		Id: id,
	}

	if err := service.Exec(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), err.Error(), c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 停止定时任务调度表
// @Description 停止定时任务调度表
// @Success 200 {object} app.Response
// @router /tools/timing/stop/:id [put]
// @Tags Admin
func (e *SysCronJobController) Stop(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt64()
	service := cron_job_service.SysCronJob{
		Id: id,
	}

	if err := service.Stop(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), err.Error(), c)
		return
	}

	response.OkWithData(nil, c)
}
