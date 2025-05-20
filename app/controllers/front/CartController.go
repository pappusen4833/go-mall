package front

import (
	"github.com/gin-gonic/gin"
	"go-mall/app/params"
	"go-mall/app/services/cart_service"
	"go-mall/pkg/app"
	"go-mall/pkg/constant"
	"go-mall/pkg/global"
	"go-mall/pkg/http/response"
	"go-mall/pkg/jwt"
	"net/http"
)

// product api
type CartController struct {
}

// @Title 购物车列表数据
// @Description 购物车列表数据
// @Success 200 {object} app.Response
// @router /api/v1/carts [get]
// @Tags Front API
func (e *CartController) CartList(c *gin.Context) {
	uid, _ := jwt.GetAppUserId(c)
	cartService := cart_service.Cart{
		Uid: uid,
	}
	vo := cartService.GetCartList()
	response.OkWithData(vo, c)

}

// @Title 获取数量
// @Description 获取数量
// @Success 200 {object} app.Response
// @router /api/v1/cart/count [get]
// @Tags Front API
func (e *CartController) Count(c *gin.Context) {
	uid, _ := jwt.GetAppUserId(c)
	cartService := cart_service.Cart{
		Uid: uid,
	}
	count := cartService.GetUserCartNum()

	response.OkWithData(gin.H{"count": count}, c)

}

// @Title 添加购物车
// @Description 添加购物车
// @Success 200 {object} app.Response
// @router /api/v1/cart/add [post]
// @Tags Front API
func (e *CartController) AddCart(c *gin.Context) {
	var (
		param params.CartParam
	)
	paramErr := app.BindAndValidate(c, &param)
	if paramErr != nil {
		response.Error(http.StatusBadRequest, 9999, paramErr.Error(), nil, c)
		return
	}
	global.LOG.Info(param)
	uid, _ := jwt.GetAppUserId(c)
	cartService := cart_service.Cart{
		Param: &param,
		Uid:   uid,
	}
	id, err := cartService.AddCart()
	if err != nil {
		response.Error(http.StatusInternalServerError, 9999, err.Error(), nil, c)
		return
	}
	response.OkWithData(gin.H{"cartId": id}, c)

}

// @Title 修改购物车数量
// @Description 修改购物车数量
// @Success 200 {object} app.Response
// @router /api/v1/cart/num [post]
// @Tags Front API
func (e *CartController) CartNum(c *gin.Context) {
	var (
		param params.CartNumParam
	)
	paramErr := app.BindAndValidate(c, &param)
	if paramErr != nil {
		response.Error(http.StatusBadRequest, 9999, paramErr.Error(), nil, c)
		return
	}
	uid, _ := jwt.GetAppUserId(c)
	cartService := cart_service.Cart{
		Uid:      uid,
		NumParam: &param,
	}
	if err := cartService.ChangeCartNum(); err != nil {
		response.Error(http.StatusInternalServerError, 9999, err.Error(), nil, c)
		return
	}
	response.OkWithData("success", c)

}

// @Title 取消收藏
// @Description 取消收藏
// @Success 200 {object} app.Response
// @router /api/v1/cart/del [post]
// @Tags Front API
func (e *CartController) DelCart(c *gin.Context) {
	var (
		param params.CartIdsParam
	)
	paramErr := app.BindAndValidate(c, &param)
	if paramErr != nil {
		response.Error(http.StatusBadRequest, 9999, paramErr.Error(), nil, c)
		return
	}

	uid, _ := jwt.GetAppUserId(c)
	cartService := cart_service.Cart{
		Uid:      uid,
		IdsParam: &param,
	}
	if err := cartService.Del(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}
	response.OkWithData("success", c)

}
