package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-mall/app/models"
	"go-mall/app/services/user_service"
	dto2 "go-mall/app/services/user_service/dto"
	"go-mall/pkg/app"
	"go-mall/pkg/constant"
	"go-mall/pkg/http/response"
	"go-mall/pkg/jwt"
	"go-mall/pkg/logging"
	"go-mall/pkg/upload"
	"go-mall/pkg/util"
	"net/http"
)

// 用户 API
type UserController struct {
}

// @Title 用户列表
// @Description 用户列表
// @Success 200 {object} app.Response
// @router /admin/user [get]
// @Tags Admin
func (e *UserController) GetAll(c *gin.Context) {
	deptId := com.StrTo(c.DefaultQuery("deptId", "-1")).MustInt64()
	enabled := com.StrTo(c.DefaultQuery("enabled", "-1")).MustInt()
	blurry := c.DefaultQuery("blurry", "")

	userService := user_service.User{
		Username: blurry,
		DeptId:   deptId,
		Enabled:  enabled,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
	}

	vo := userService.GetUserAll()
	response.OkWithData(vo, c)
}

// @Title 用户添加
// @Description 用户添加
// @Success 200 {object} app.Response
// @router /admin/user [post]
// @Tags Admin
func (e *UserController) Post(c *gin.Context) {
	var (
		model models.SysUser
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	userService := user_service.User{
		M: &model,
	}

	if err := userService.Insert(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)

}

// @Title 用户编辑
// @Description 用户编辑
// @Success 200 {object} app.Response
// @router /admin/user [put]
// @Tags Admin
func (e *UserController) Put(c *gin.Context) {
	var (
		model models.SysUser
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	userService := user_service.User{
		M: &model,
	}

	if err := userService.Save(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 用户删除
// @Description 用户删除
// @Success 200 {object} app.Response
// @router /admin/user [delete]
// @Tags Admin
func (e *UserController) Delete(c *gin.Context) {
	var (
		ids []int64
	)
	c.BindJSON(&ids)
	userService := user_service.User{Ids: ids}

	if err := userService.Del(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)

}

// @Title 用户上传图像
// @Description 用户上传图像
// @Success 200 {object} app.Response
// @router /admin/updateAvatar [post]
// @Tags Admin
func (e *UserController) Avatar(c *gin.Context) {
	file, image, err := c.Request.FormFile("file")
	if err != nil {
		logging.Warn(err)
		response.Error(http.StatusInternalServerError, constant.ERROR, constant.GetMsg(constant.ERROR), nil, c)
		return
	}

	if image == nil {
		response.Error(http.StatusBadRequest, constant.INVALID_PARAMS, constant.GetMsg(constant.INVALID_PARAMS), nil, c)
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		response.Error(http.StatusBadRequest, constant.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, constant.GetMsg(constant.ERROR_UPLOAD_CHECK_IMAGE_FORMAT), nil, c)
		return
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		logging.Warn(err)
		response.Error(http.StatusInternalServerError, constant.ERROR_UPLOAD_CHECK_IMAGE_FAIL, constant.GetMsg(constant.ERROR_UPLOAD_CHECK_IMAGE_FAIL), nil, c)
		return
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		logging.Warn(err)
		response.Error(http.StatusInternalServerError, constant.ERROR_UPLOAD_SAVE_IMAGE_FAIL, constant.GetMsg(constant.ERROR_UPLOAD_SAVE_IMAGE_FAIL), nil, c)
		return
	}

	uid, _ := jwt.GetAdminUserId(c)
	userService := user_service.User{ImageUrl: upload.GetImageFullUrl(imageName), Id: uid}
	if err := userService.UpdateImage(); err != nil {
		response.Error(http.StatusInternalServerError, 9999, err.Error(), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 用户修改密码
// @Description 用户修改密码
// @Success 200 {object} app.Response
// @router /admin/updatePass [post]
// @Tags Admin
func (e *UserController) Pass(c *gin.Context) {
	var (
		model dto2.UserPass
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	uid, _ := jwt.GetAdminUserId(c)
	userService := user_service.User{UserPass: model, Id: uid}
	if err := userService.UpdatePass(); err != nil {
		response.Error(http.StatusInternalServerError, 9999, err.Error(), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 用户修改个人信息
// @Description 用户修改个人信息
// @Success 200 {object} app.Response
// @router /admin/center [put]
// @Tags Admin
func (e *UserController) Center(c *gin.Context) {
	var (
		model dto2.UserPost
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	uid, _ := jwt.GetAdminUserId(c)
	userService := user_service.User{UserPost: model, Id: uid}
	if err := userService.UpdateProfile(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Summary Import Image
// @Produce  json
// @Param image formData file true "Image File"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @router /admin/upload [post]
// @Tags Admin
func UploadImage(c *gin.Context) {
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		response.Error(http.StatusInternalServerError, constant.ERROR, constant.GetMsg(constant.ERROR), nil, c)
		return
	}

	if image == nil {
		response.Error(http.StatusBadRequest, constant.INVALID_PARAMS, constant.GetMsg(constant.INVALID_PARAMS), nil, c)
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		response.Error(http.StatusBadRequest, constant.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, constant.GetMsg(constant.ERROR_UPLOAD_CHECK_IMAGE_FORMAT), nil, c)
		return
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		logging.Warn(err)
		response.Error(http.StatusInternalServerError, constant.ERROR_UPLOAD_CHECK_IMAGE_FAIL, constant.GetMsg(constant.ERROR_UPLOAD_CHECK_IMAGE_FAIL), nil, c)
		return
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		logging.Warn(err)
		response.Error(http.StatusInternalServerError, constant.ERROR_UPLOAD_SAVE_IMAGE_FAIL, constant.GetMsg(constant.ERROR_UPLOAD_SAVE_IMAGE_FAIL), nil, c)
		return
	}

	response.OkWithData(map[string]string{
		"image_url":      upload.GetImageFullUrl(imageName),
		"image_save_url": savePath + imageName,
	}, c)
}
