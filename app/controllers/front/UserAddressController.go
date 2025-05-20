package front

import (
	"github.com/gin-gonic/gin"
	"go-mall/app/params"
	"go-mall/app/services/address_service"
	"go-mall/pkg/app"
	"go-mall/pkg/http/response"
	"go-mall/pkg/jwt"
	"go-mall/pkg/util"
	"net/http"
)

// Address api
type UserAddressController struct {
}

// @Title 设置默认地址
// @Description 设置默认地址
// @Success 200 {object} app.Response
// @router /api/v1/address/del [post]
// @Tags Front API
func (e *UserAddressController) Del(c *gin.Context) {
	var (
		param params.IdParam
	)
	paramErr := app.BindAndValidate(c, &param)
	if paramErr != nil {
		response.Error(http.StatusBadRequest, 9999, paramErr.Error(), nil, c)
		return
	}

	uid, _ := jwt.GetAppUserId(c)
	addressService := address_service.Address{
		Id:  param.Id,
		Uid: uid,
	}
	err := addressService.DelAddress()
	if err != nil {
		response.Error(http.StatusInternalServerError, 9999, err.Error(), nil, c)
		return
	}
	response.OkWithData("ok", c)

}

// @Title 设置默认地址
// @Description 设置默认地址
// @Success 200 {object} app.Response
// @router /api/v1/address/default/set [post]
// @Tags Front API
func (e *UserAddressController) SetDefault(c *gin.Context) {
	var (
		param params.IdParam
	)
	paramErr := app.BindAndValidate(c, &param)
	if paramErr != nil {
		response.Error(http.StatusBadRequest, 9999, paramErr.Error(), nil, c)
		return
	}

	uid, _ := jwt.GetAppUserId(c)
	addressService := address_service.Address{
		Id:  param.Id,
		Uid: uid,
	}
	err := addressService.SetDefault()
	if err != nil {
		response.Error(http.StatusInternalServerError, 9999, err.Error(), nil, c)
		return
	}
	response.OkWithData("ok", c)

}

// @Title 获取列表数据
// @Description 获取列表数据
// @Success 200 {object} app.Response
// @router /api/v1/address [get]
// @Tags Front API
func (e *UserAddressController) GetList(c *gin.Context) {
	uid, _ := jwt.GetAppUserId(c)
	addressService := address_service.Address{
		Enabled:  1,
		PageNum:  util.GetFrontPage(c),
		PageSize: util.GetFrontLimit(c),
		Uid:      uid,
	}

	vo, total, page := addressService.GetList()

	response.PageResult(0, vo, "ok", total, page, c)
}

// @Title 添加or更新地址
// @Description 添加or更新地址
// @Success 200 {object} app.Response
// @router /api/v1/address/edit [post]
// @Tags Front API
func (e *UserAddressController) SaveAddress(c *gin.Context) {
	var (
		param params.AddressParan
	)
	paramErr := app.BindAndValidate(c, &param)
	if paramErr != nil {
		response.Error(http.StatusBadRequest, 9999, paramErr.Error(), nil, c)
		return
	}

	uid, _ := jwt.GetAppUserId(c)
	addressService := address_service.Address{
		Param: &param,
		Uid:   uid,
	}
	id, err := addressService.AddOrUpdate()
	if err != nil {
		response.Error(http.StatusInternalServerError, 9999, err.Error(), nil, c)
		return
	}
	response.OkWithData(gin.H{"id": id}, c)

}

// @Title 获取树形数据
// @Description 获取树形数据
// @Success 200 {object} app.Response
// @router /api/v1/city_list [get]
// @Tags Front API
func (e *UserAddressController) GetCityList(c *gin.Context) {
	addressService := address_service.Address{Enabled: 1}
	vo := addressService.GetCitys()
	response.OkWithData(vo, c)

}
