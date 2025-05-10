package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-mall/app/models"
	"go-mall/app/services/canvas_service"
	"go-mall/pkg/app"
	"go-mall/pkg/constant"
	"go-mall/pkg/http/response"
	"net/http"
)

// 画布api
type CanvasController struct {
}

// @Title 画布
// @Description 画布
// @Success 200 {object} app.Response
// @router /admin/canvas/getCanvas [get]
// @Tags Admin
func (e *CanvasController) Get(c *gin.Context) {
	terminal := com.StrTo(c.DefaultQuery("terminal", "3")).MustInt()
	canvasService := canvas_service.Canvas{
		Terminal: terminal,
	}
	vo := canvasService.Get()
	response.OkWithData(vo, c)
}

// @Title 画布添加/修改
// @Description 画布添加/修改
// @Success 200 {object} app.Response
// @router /admin/canvas/saveCanvas [post]
// @Tags Admin
func (e *CanvasController) Post(c *gin.Context) {
	var (
		model models.StoreCanvas
	)
	paramErr := app.BindAndValidate(c, &model)
	if paramErr != nil {
		response.Error(http.StatusBadRequest, 9999, paramErr.Error(), nil, c)
		return
	}

	canvasService := canvas_service.Canvas{
		M: &model,
	}

	if err := canvasService.Save(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)

}
