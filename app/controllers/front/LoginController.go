package front

import (
	"github.com/gin-gonic/gin"
	"go-mall/app/params"
	"go-mall/app/services/wechat_user_service"
	"go-mall/pkg/app"
	"go-mall/pkg/constant"
	"go-mall/pkg/http/response"
	"go-mall/pkg/jwt"
	"go-mall/pkg/util"
	"net/http"
	"time"
)

// 登录api
type LoginController struct {
}

// @Title 会员登录
// @Description 会员登录
// @Success 200 {object} app.Response
// @router /api/v1/login [post]
// @Param data body string true "body data"
// @Tags Front API
func (e *LoginController) Login(c *gin.Context) {
	var (
		param params.HLoginParam
	)
	paramErr := app.BindAndValidate(c, &param)
	if paramErr != nil {
		response.Error(http.StatusBadRequest, 9999, paramErr.Error(), nil, c)
		return
	}
	userService := wechat_user_service.User{HLoginParam: &param}
	user, err := userService.HLogin()
	if err != nil {
		response.Error(http.StatusInternalServerError, 9999, err.Error(), nil, c)
		return
	}

	d := time.Now().Add(time.Hour * 24 * 100)
	token, _ := jwt.GenerateAppToken(user, d)
	response.OkWithData(
		gin.H{
			"token":        token,
			"expires_time": d.Unix(),
		}, c)
}

func (e *LoginController) AutoRegisterAndLogin(c *gin.Context) {
	var (
		param  params.RegParam
		hparam params.HLoginParam
	)

	param.Account, param.Password = util.RandomString(10), util.RandomString(32)
	hparam.Username, hparam.Password = param.Account, param.Password
	userService := wechat_user_service.User{RegParam: &param, HLoginParam: &hparam, Ip: c.ClientIP()}

	if err := userService.Reg(); err != nil {
		response.Error(http.StatusInternalServerError, 9999, err.Error(), nil, c)
		return
	}

	user, err := userService.HLogin()
	if err != nil {
		response.Error(http.StatusInternalServerError, 9999, err.Error(), nil, c)
		return
	}

	d := time.Now().Add(time.Hour * 24 * 100)
	token, _ := jwt.GenerateAppToken(user, d)
	response.OkWithData(
		gin.H{
			"token":        token,
			"expires_time": d.Unix(),
		}, c)
}

// @Title 短信验证码
// @Description 短信验证码
// @Success 200 {object} app.Response
// @router /api/v1/register/verify [post]
// @Tags Front API
func (e *LoginController) Verify(c *gin.Context) {
	var (
		param params.VerityParam
	)
	paramErr := app.BindAndValidate(c, &param)
	if paramErr != nil {
		response.Error(http.StatusBadRequest, 9999, paramErr.Error(), nil, c)
		return
	}
	userService := wechat_user_service.User{VerityParam: &param}
	str, err := userService.Verify()
	if err != nil {
		response.Error(http.StatusInternalServerError, 9999, err.Error(), nil, c)
		return
	}

	response.OkWithData(str, c)

}

// @Title 注册
// @Description 注册
// @Success 200 {object} app.Response
// @router /api/v1/register [post]
// @Tags Front API
func (e *LoginController) Reg(c *gin.Context) {
	var (
		param params.RegParam
	)
	paramErr := app.BindAndValidate(c, &param)
	if paramErr != nil {
		response.Error(http.StatusBadRequest, 9999, paramErr.Error(), nil, c)
		return
	}
	userService := wechat_user_service.User{RegParam: &param, Ip: c.ClientIP()}
	if err := userService.Reg(); err != nil {
		response.Error(http.StatusInternalServerError, 9999, err.Error(), nil, c)
		return
	}

	response.OkWithData("success", c)

}

// @Title 获取用户信息
// @Description 获取用户信息
// @Success 200 {object} app.Response
// @router /api/v1/info [get]
// @Tags Front API
func (e *LoginController) Info(c *gin.Context) {
	response.OkWithData(jwt.GetAdminDetailUser(c), c)
}

// @Title 退出登录
// @Description 退出登录
// @Success 200 {object} app.Response
// @router /api/v1/logout [delete]
// @Tags Front API
func (e *LoginController) Logout(c *gin.Context) {
	err := jwt.RemoveUser(c)
	if err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_LOGOUT_USER, constant.GetMsg(constant.FAIL_LOGOUT_USER), nil, c)
		return
	}

	response.OkWithData(nil, c)
}
