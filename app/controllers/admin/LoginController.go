package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go-mall/app/models/dto"
	"go-mall/app/models/vo"
	"go-mall/app/services/user_service"
	"go-mall/pkg/app"
	"go-mall/pkg/constant"
	"go-mall/pkg/http/response"
	"go-mall/pkg/jwt"
	"go-mall/pkg/logging"
	"go-mall/pkg/util"
	"image/color"
	"net/http"
	"time"
)

// 登录api
type LoginController struct {
}

type CaptchaResult struct {
	Id          string `json:"id"`
	Base64Blob  string `json:"base_64_blob"`
	VerifyValue string `json:"code"`
}

// 设置自带的store
var store = base64Captcha.DefaultMemStore

// @Title 登录
// @Description 登录
// @Success 200 {object} app.Response
// @router /admin/login [post]
// @Tags Admin
func (e *LoginController) Login(c *gin.Context) {
	var (
		authUser dto.AuthUser
	)

	//body, _ := ioutil.ReadAll(c.Request.Body)
	//logging.Info(string(body))

	httpCode, errCode := app.BindAndValid(c, &authUser)
	logging.Info(authUser)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}

	userService := user_service.User{Username: authUser.Username}
	currentUser, err := userService.GetUserOneByName()
	if err != nil {
		response.Error(http.StatusInternalServerError, constant.ERROR_NOT_EXIST_USER, constant.GetMsg(constant.ERROR_NOT_EXIST_USER), nil, c)
		return
	}

	//校验验证码
	if !store.Verify(authUser.Id, authUser.Code, true) {
		response.Error(http.StatusInternalServerError, constant.ERROR_CAPTCHA_USER, constant.GetMsg(constant.ERROR_CAPTCHA_USER), nil, c)
		return
	}
	if !util.ComparePwd(currentUser.Password, []byte(authUser.Password)) {
		response.Error(http.StatusInternalServerError, constant.ERROR_PASS_USER, constant.GetMsg(constant.ERROR_PASS_USER), nil, c)
		return
	}
	token, _ := jwt.GenerateToken(currentUser, time.Hour*24*100)
	var loginVO = new(vo.LoginVo)
	loginVO.Token = token
	loginVO.User = currentUser
	response.OkWithData(loginVO, c)

}

// @Title 获取用户信息
// @Description 获取用户信息
// @Success 200 {object} app.Response
// @router /admin/info [get]
// @Tags Admin
func (e *LoginController) Info(c *gin.Context) {
	response.OkWithData(jwt.GetAdminDetailUser(c), c)
}

// @Title 退出登录
// @Description 退出登录
// @Success 200 {object} app.Response
// @router /admin/logout [delete]
// @Tags Admin
func (e *LoginController) Logout(c *gin.Context) {
	err := jwt.RemoveUser(c)
	if err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_LOGOUT_USER, constant.GetMsg(constant.FAIL_LOGOUT_USER), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 获取验证码
// @Description 获取验证码
// @router /admin/captcha [get]
// @Tags Admin
func (e *LoginController) Captcha(c *gin.Context) {
	GenerateCaptcha(c)
}

// 生成图形化验证码  ctx *context.Context
func GenerateCaptcha(c *gin.Context) {
	var (
		driver       base64Captcha.Driver
		driverString base64Captcha.DriverMath
	)

	// 配置验证码信息
	captchaConfig := base64Captcha.DriverMath{
		Height:          38,
		Width:           110,
		NoiseCount:      0,
		ShowLineOptions: 0,
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	// 自定义配置，如果不需要自定义配置，则上面的结构体和下面这行代码不用写
	driverString = captchaConfig
	driver = driverString.ConvertFonts()

	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()
	if err != nil {
		logging.Error(err.Error())
	}
	captchaResult := CaptchaResult{
		Id:         id,
		Base64Blob: b64s,
	}

	response.OkWithData(captchaResult, c)
}
