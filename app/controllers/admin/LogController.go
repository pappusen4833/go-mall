package admin

import (
	"github.com/gin-gonic/gin"
	"go-mall/app/services/log_service"
	"go-mall/pkg/constant"
	"go-mall/pkg/http/response"
	"go-mall/pkg/util"
	"net/http"
)

// 角色 API
type LogController struct {
}

// @Title 日志列表
// @Description 日志列表
// @Success 200 {object} app.Response
// @router /admin/logs [get]
// @Tags Admin
func (e *LogController) GetAll(c *gin.Context) {
	blurry := c.DefaultQuery("blurry", "")
	logService := log_service.Log{
		Des:      blurry,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
	}
	vo := logService.GetAll()
	response.OkWithData(vo, c)
}

// @Title 日志删除
// @Description 日志删除
// @Success 200 {object} app.Response
// @router /admin/logs [delete]
// @Tags Admin
func (e *LogController) Delete(c *gin.Context) {
	var (
		ids []int64
	)
	c.BindJSON(&ids)
	logService := log_service.Log{Ids: ids}

	if err := logService.Del(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}
