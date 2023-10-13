/**
* Copyright (C) 2020-2021
* All rights reserved, Designed By www.yixiang.co
* 注意：本软件为www.yixiang.co开发研制
 */
package front

import (
	"github.com/gin-gonic/gin"
	"go-mall/app/services/product_relation_service"
	"go-mall/app/services/wechat_user_service"
	"go-mall/packages/app"
	"go-mall/packages/constant"
	"go-mall/packages/jwt"
	"go-mall/packages/util"
	"net/http"
)

// user api
type UserController struct {
}

// @Title 获取用户信息
// @Description 获取用户信息
// @Success 200 {object} app.Response
// @router /api/v1/userinfo [get]
func (e *UserController) GetUserInfo(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	user, _ := jwt.GetAppDetailUser(c)
	userService := wechat_user_service.User{User: user}
	vo := userService.GetUserDetail()
	appG.Response(http.StatusOK, constant.SUCCESS, vo)

}

// @Title 获取用户收藏
// @Description 获取用户收藏
// @Success 200 {object} app.Response
// @router /api/v1/collect/user [get]
func (e *UserController) CollectUser(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	uid, _ := jwt.GetAppUserId(c)
	relationService := product_relation_service.Relation{
		PageNum:  util.GetFrontPage(c),
		PageSize: util.GetFrontLimit(c),
		Uid:      uid,
	}
	vo, total, page := relationService.GetUserCollectList()
	appG.ResponsePage(http.StatusOK, constant.SUCCESS, vo, total, page)

}
