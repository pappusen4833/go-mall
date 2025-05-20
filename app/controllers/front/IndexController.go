package front

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-mall/app/services/canvas_service"
	"go-mall/app/services/product_service"
	"go-mall/pkg/constant"
	productEnum "go-mall/pkg/enums/product"
	"go-mall/pkg/http/response"
	"go-mall/pkg/logging"
	"go-mall/pkg/upload"
	"net/http"
)

// index api
type IndexController struct {
}

// @Title 获取首页数据
// @Description 获取首页数据
// @Success 200 {object} app.Response
// @router /api/v1/index [get]
// @Tags Front API
func (e *IndexController) GetIndex(c *gin.Context) {
	productService := product_service.Product{
		Enabled:  1,
		PageNum:  0,
		PageSize: 6,
		Order:    productEnum.STATUS_1,
	}

	vo1, _, _ := productService.GetList()

	productService.PageSize = 10
	productService.Order = productEnum.STATUS_2
	vo2, _, _ := productService.GetList()

	productService.PageSize = 6
	productService.Order = productEnum.STATUS_3
	vo3, _, _ := productService.GetList()

	productService.PageSize = 10
	productService.Order = productEnum.STATUS_4
	vo4, _, _ := productService.GetList()
	res := gin.H{
		"bastList":  vo1,
		"likeInfo":  vo2,
		"firstList": vo3,
		"benefit":   vo4,
	}
	response.OkWithData(res, c)

}

// @Title 获取画布数据
// @Description 获取画布数据
// @Success 200 {object} app.Response
// @router /api/v1/getCanvas [get]
// @Tags Front API
func (e *IndexController) GetCanvas(c *gin.Context) {
	terminal := com.StrTo(c.DefaultQuery("terminal", "3")).MustInt()
	canvasService := canvas_service.Canvas{
		Terminal: terminal,
	}
	vo := canvasService.Get()
	response.OkWithData(vo, c)

}

// @Title 上传图像
// @Description 上传图像
// @Success 200 {object} app.Response
// @router /upload [post]
// @Tags Front API
func (e *IndexController) Upload(c *gin.Context) {
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
