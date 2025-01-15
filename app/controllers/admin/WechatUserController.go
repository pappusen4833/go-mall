package admin

import (
	"github.com/gin-gonic/gin"
	dto2 "go-mall/app/services/user_service/dto"
	"go-mall/app/services/wechat_user_service"
	dto3 "go-mall/app/services/wechat_user_service/dto"
	"go-mall/pkg/app"
	"go-mall/pkg/constant"
	"go-mall/pkg/http/response"
	"go-mall/pkg/util"
	"net/http"
)

// 微信用户 API
type WechatUserController struct {
}

// @Title 用户列表
// @Description 用户列表
// @Success 200 {object} app.Response
// @router /weixin/user [get]
// @Tags Admin
func (e *WechatUserController) GetAll(c *gin.Context) {
	value := c.DefaultQuery("value", "")
	myType := c.DefaultQuery("type", "")
	userType := c.DefaultQuery("userType", "")

	userService := wechat_user_service.User{
		Value:    value,
		MyType:   myType,
		UserType: userType,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
	}

	vo := userService.GetUserAll()
	response.OkWithData(vo, c)
}

// @Title 用户编辑
// @Description 用户编辑
// @Success 200 {object} app.Response
// @router /weixin/user [put]
// @Tags Admin
func (e *WechatUserController) Put(c *gin.Context) {
	var (
		model dto3.User
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	userService := wechat_user_service.User{
		Dto: &model,
	}

	if err := userService.Save(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 用户余额修改
// @Description 用户余额修改
// @Success 200 {object} app.Response
// @router /weixin/user/money [post]
// @Tags Admin
func (e *WechatUserController) Money(c *gin.Context) {
	var (
		model dto2.UserMoney
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	userService := wechat_user_service.User{
		Money: &model,
	}

	if err := userService.SaveMoney(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}
