package admin

import (
	"github.com/gin-gonic/gin"
	"go-mall/app/services/wechat_menu_service"
	dto2 "go-mall/app/services/wechat_menu_service/dto"
	"go-mall/pkg/app"
	"go-mall/pkg/constant"
	"go-mall/pkg/http/response"
	"net/http"
)

// 菜单api
type WechatMenuController struct {
}

// @Title 获取菜单
// @Description 获取菜单
// @Success 200 {object} app.Response
// @router /weixin/menu [get]
// @Tags Admin
func (e *WechatMenuController) GetAll(c *gin.Context) {
	meuService := wechat_menu_service.Menu{}
	vo := meuService.GetAll()
	response.OkWithData(vo, c)
}

// @Title 菜单更新
// @Description 菜单更新
// @Success 200 {object} app.Response
// @router /weixin/menu [post]
// @Tags Admin
func (e *WechatMenuController) Post(c *gin.Context) {
	var (
		dto dto2.WechatMenu
	)
	httpCode, errCode := app.BindAndValid(c, &dto)
	if errCode != constant.SUCCESS {
		response.Error(httpCode, errCode, constant.GetMsg(errCode), nil, c)
		return
	}
	meuService := wechat_menu_service.Menu{
		Dto: dto,
	}

	if err := meuService.Insert(); err != nil {
		response.Error(http.StatusInternalServerError, constant.FAIL_ADD_DATA, constant.GetMsg(constant.FAIL_ADD_DATA), nil, c)
		return
	}

	response.OkWithData(nil, c)

}
