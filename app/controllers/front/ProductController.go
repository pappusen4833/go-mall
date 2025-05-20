package front

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-mall/app/params"
	"go-mall/app/services/product_relation_service"
	"go-mall/app/services/product_reply_service"
	"go-mall/app/services/product_service"
	"go-mall/pkg/app"
	productEnum "go-mall/pkg/enums/product"
	"go-mall/pkg/http/response"
	"go-mall/pkg/jwt"
	"go-mall/pkg/util"
	"net/http"
)

// product api
type ProductController struct {
}

// @Title 获取商品列表数据
// @Description 获取商品列表数据
// @Success 200 {object} app.Response
// @router /api/v1/products [get]
// @Tags Front API
func (e *ProductController) GoodsList(c *gin.Context) {
	productService := product_service.Product{
		Name:       c.Query("keyword"),
		Enabled:    1,
		PageNum:    util.GetFrontPage(c),
		PageSize:   util.GetFrontLimit(c),
		Sid:        c.Query("sid"),
		News:       c.Query("news"),
		PriceOrder: c.Query("priceOrder"),
		SalesOrder: c.Query("salesOrder"),
	}

	vo, total, page := productService.GetList()

	response.PageResult(0, vo, "ok", total, page, c)
}

// @Title 获取推荐商品
// @Description 获取推荐商品
// @Success 200 {object} app.Response
// @router /api/v1/product/hot [get]
// @Tags Front API
func (e *ProductController) GoodsRecommendList(c *gin.Context) {
	productService := product_service.Product{
		Enabled:  1,
		PageNum:  0,
		PageSize: 6,
		Order:    productEnum.STATUS_1,
	}

	vo, _, _ := productService.GetList()

	response.OkWithData(vo, c)

}

// @Title 获取商品详情
// @Description 获取商品详情
// @Success 200 {object} app.Response
// @router /api/v1/product/detail/:id [get]
// @Tags Front API
func (e *ProductController) GoodDetail(c *gin.Context) {
	var (
		uid int64
	)
	id := com.StrTo(c.Param("id")).MustInt64()
	user, err := jwt.GetAppDetailUser(c)
	if err != nil {
		uid = 0
	} else {
		uid = user.Id
	}

	productService := product_service.Product{
		Id:  id,
		Uid: uid,
	}

	vo, err := productService.GetDetail()
	if err != nil {
		response.Error(http.StatusBadRequest, 9999, err.Error(), nil, c)
		return
	}

	response.OkWithData(vo, c)
}

// @Title 添加收藏
// @Description 添加收藏
// @Success 200 {object} app.Response
// @router /api/v1/collect/add [post]
// @Tags Front API
func (e *ProductController) AddCollect(c *gin.Context) {
	var (
		param params.RelationParam
	)
	paramErr := app.BindAndValidate(c, &param)
	if paramErr != nil {
		response.Error(http.StatusBadRequest, 9999, paramErr.Error(), nil, c)
		return
	}
	uid, _ := jwt.GetAppUserId(c)
	relationService := product_relation_service.Relation{
		Param: &param,
		Uid:   uid,
	}
	if err := relationService.AddRelation(); err != nil {
		response.Error(http.StatusInternalServerError, 9999, err.Error(), nil, c)
		return
	}
	response.OkWithData("success", c)

}

// @Title 取消收藏
// @Description 取消收藏
// @Success 200 {object} app.Response
// @router /api/v1/collect/del [post]
// @Tags Front API
func (e *ProductController) DelCollect(c *gin.Context) {
	var (
		param params.RelationParam
	)
	paramErr := app.BindAndValidate(c, &param)
	if paramErr != nil {
		response.Error(http.StatusBadRequest, 9999, paramErr.Error(), nil, c)
		return
	}
	uid, _ := jwt.GetAppUserId(c)
	relationService := product_relation_service.Relation{
		Param: &param,
		Uid:   uid,
	}
	if err := relationService.DelRelation(); err != nil {
		response.Error(http.StatusInternalServerError, 9999, err.Error(), nil, c)
		return
	}
	response.OkWithData("success", c)

}

// @Title 获取商品评论列表数据
// @Description 获取商品评论列表数据
// @Success 200 {object} app.Response
// @router /api/v1/reply/list/:id [get]
// @Tags Front API
func (e *ProductController) ReplyList(c *gin.Context) {
	replyService := product_reply_service.Reply{
		ProductId: com.StrTo(c.Param("id")).MustInt64(),
		PageNum:   util.GetFrontPage(c),
		PageSize:  util.GetFrontLimit(c),
		Type:      com.StrTo(c.Query("type")).MustInt(),
	}

	vo, total, page := replyService.GetList()

	response.PageResult(0, vo, "ok", total, page, c)
}
