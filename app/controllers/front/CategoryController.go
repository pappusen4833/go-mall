package front

import (
	"github.com/gin-gonic/gin"
	"go-mall/app/services/cate_service"
	"go-mall/pkg/http/response"
)

// category api
type CategoryController struct {
}

// @Title 获取树形数据
// @Description 获取树形数据
// @Success 200 {object} app.Response
// @router /api/v1/category [get]
// @Tags Front API
func (e *CategoryController) GetCateList(c *gin.Context) {
	cateService := cate_service.Cate{Enabled: 1}
	vo := cateService.GetAll()
	response.OkWithData(vo, c)

}
