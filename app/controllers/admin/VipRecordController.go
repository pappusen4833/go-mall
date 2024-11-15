package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-mall/app/models"
	"go-mall/app/services/vip_record_service"
	"go-mall/pkg/app"
	"go-mall/pkg/constant"
	"go-mall/pkg/http/response"
	"go-mall/pkg/util"
	"net/http"
)

type VipRecordController struct {
}

// @Title 获取用户VIP列表
// @Description 获取用户VIP列表
// @Success 200 {object} app.Response
// @router / [get]
func (e *VipRecordController) GetAll(c *gin.Context) {
	enabled := com.StrTo(c.DefaultQuery("enabled", "-1")).MustInt()
	name := c.DefaultQuery("blurry", "")
	service := vip_record_service.VipRecord{
		Enabled:  enabled,
		Name:     name,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
	}
	vo := service.GetAll()
	response.OkWithData(vo, c)
}

// @Title 添加用户VIP
// @Description 添加用户VIP
// @Success 200 {object} app.Response
// @router / [post]
func (e *VipRecordController) Post(c *gin.Context) {
	var (
		model models.VipRecord
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	service := vip_record_service.VipRecord{
		M: &model,
	}

	if err := service.Insert(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 修改用户VIP
// @Description 修改用户VIP
// @Success 200 {object} app.Response
// @router / [put]
func (e *VipRecordController) Put(c *gin.Context) {
	var (
		model models.VipRecord
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	service := vip_record_service.VipRecord{
		M: &model,
	}

	if err := service.Save(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 删除用户VIP
// @Description 删除用户VIP
// @Success 200 {object} app.Response
// @router /:id [delete]
func (e *VipRecordController) Delete(c *gin.Context) {
	var (
		ids []int64
	)
	if strId := c.Param("id"); strId != "" {
		id := com.StrTo(strId).MustInt64()
		ids = append(ids, id)
	} else {
		c.BindJSON(&ids)
	}

	service := vip_record_service.VipRecord{Ids: ids}
	if err := service.Del(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}
