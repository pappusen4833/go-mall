package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-mall/app/models"
	"go-mall/app/services/material_service"
	"go-mall/pkg/app"
	"go-mall/pkg/constant"
	"go-mall/pkg/global"
	"go-mall/pkg/http/response"
	"go-mall/pkg/jwt"
	"go-mall/pkg/logging"
	"go-mall/pkg/upload"
	"go-mall/pkg/util"
	"net/http"
)

// 素材api
type MaterialController struct {
}

// @Title 素材列表
// @Description 岗位列表
// @Success 200 {object} app.Response
// @router /admin/material [get]
// @Tags Admin
func (e *MaterialController) GetAll(c *gin.Context) {
	groupId := com.StrTo(c.DefaultQuery("groupId", "-1")).MustInt64()
	name := c.DefaultQuery("blurry", "")
	materialService := material_service.Material{
		GroupId:  groupId,
		Name:     name,
		PageSize: util.GetSize(c),
		PageNum:  util.GetPage(c),
	}
	vo := materialService.GetAll()
	response.OkWithData(vo, c)
}

// @Title 素材添加
// @Description 素材添加
// @Success 200 {object} app.Response
// @router /admin/material [post]
// @Tags Admin
func (e *MaterialController) Post(c *gin.Context) {
	var (
		model models.SysMaterial
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	uid, _ := jwt.GetAdminUserId(c)
	model.CreateId = uid
	materialService := material_service.Material{
		M: &model,
	}

	if err := materialService.Insert(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 素材修改
// @Description 素材修改
// @Success 200 {object} app.Response
// @router /admin/material [put]
// @Tags Admin
func (e *MaterialController) Put(c *gin.Context) {
	var (
		model models.SysMaterial
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	uid, _ := jwt.GetAdminUserId(c)
	model.CreateId = uid
	materialService := material_service.Material{
		M: &model,
	}

	if err := materialService.Save(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 素材删除
// @Description 素材删除
// @Success 200 {object} app.Response
// @router /admin/material/:id [delete]
// @Tags Admin
func (e *MaterialController) Delete(c *gin.Context) {
	var (
		ids []int64
	)
	id := com.StrTo(c.Param("id")).MustInt64()
	ids = append(ids, id)
	//c.BindJSON(&ids)
	materialService := material_service.Material{Ids: ids}

	if err := materialService.Del(); err != nil {
		global.LOG.Error(err)
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 上传图像
// @Description 上传图像
// @Success 200 {object} app.Response
// @router /admin/material/upload [post]
// @Tags Admin
func (e *MaterialController) Upload(c *gin.Context) {
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
	//savePath := upload.GetImagePath()
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

	imageUrl := upload.GetImageFullUrl(imageName)
	//imageSaveUrl := avePath + imageName

	response.OkWithData(imageUrl, c)

}
