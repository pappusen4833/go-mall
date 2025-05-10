package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-mall/app/models"
	"go-mall/app/services/cate_service"
	"go-mall/pkg/app"
	"go-mall/pkg/constant"
	"go-mall/pkg/http/response"
	"net/http"
)

// 商品分类api
type StoreCategoryController struct {
}

// @Title 商品分类列表
// @Description 商品分类列表
// @Success 200 {object} app.Response
// @router /shop/cate [get]
// @Tags Admin
func (e *StoreCategoryController) GetAll(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	enabled := com.StrTo(c.DefaultQuery("enabled", "-1")).MustInt()
	cateService := cate_service.Cate{Name: name, Enabled: enabled}
	vo := cateService.GetAll()
	response.OkWithData(vo, c)
}

// @Title 添加商品分类
// @Description 添加商品分类
// @Success 200 {object} app.Response
// @router /shop/cate [post]
// @Tags Admin
func (e *StoreCategoryController) Post(c *gin.Context) {
	var (
		model models.StoreCategory
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	cateService := cate_service.Cate{
		M: &model,
	}

	if err := cateService.Insert(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 修改商品分类
// @Description 修改商品分类
// @Success 200 {object} app.Response
// @router /shop/cate [put]
// @Tags Admin
func (e *StoreCategoryController) Put(c *gin.Context) {
	var (
		model models.StoreCategory
	)
	httpCode, errCode := app.BindAndValid(c, &model)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	cateService := cate_service.Cate{
		M: &model,
	}

	if err := cateService.Save(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}

// @Title 删除商品分类
// @Description 删除商品分类
// @Success 200 {object} app.Response
// @router /shop/cate [delete]
// @Tags Admin
func (e *StoreCategoryController) Delete(c *gin.Context) {
	var (
		ids []int64
	)
	c.BindJSON(&ids)
	cateService := cate_service.Cate{Ids: ids}

	if err := cateService.Del(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)
}
