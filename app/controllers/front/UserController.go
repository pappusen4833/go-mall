package front

import (
	"github.com/gin-gonic/gin"
	"go-mall/app/models"
	"go-mall/app/services/product_relation_service"
	"go-mall/app/services/wechat_user_service"
	"go-mall/pkg/global"
	"go-mall/pkg/http/response"
	"go-mall/pkg/jwt"
	"go-mall/pkg/util"
)

// user api
type UserController struct {
}

// GetUserInfo @Title 获取用户信息
// @Description 获取用户信息
// @Success 200 {object} app.Response
// @router /api/v1/userinfo [get]
// @Tags Front API
func (e *UserController) GetUserInfo(c *gin.Context) {
	user, _ := jwt.GetAppDetailUser(c)
	userService := wechat_user_service.User{User: user}
	vo := userService.GetUserDetail()
	response.OkWithData(vo, c)

}

func (e *UserController) Info(c *gin.Context) {
	user, _ := jwt.GetAppDetailUser(c)
	global.LOG.Info("jwt user detail")
	global.LOG.Info(user)
	// get member record
	vipRecord := models.FindVipRecordByUserId(user.Id)
	userInfo := make(map[string]interface{})
	userInfo["user_id"] = user.Id
	userInfo["nickname"] = user.Nickname
	userInfo["headimgurl"] = user.Avatar
	userInfo["type"] = 1
	userInfo["vip"] = 0
	userInfo["vipRecord"] = vipRecord
	response.OkWithData(userInfo, c)
}

// CollectUser @Title 获取用户收藏
// @Description 获取用户收藏
// @Success 200 {object} app.Response
// @router /api/v1/collect/user [get]
// @Tags Front API
func (e *UserController) CollectUser(c *gin.Context) {
	uid, _ := jwt.GetAppUserId(c)
	relationService := product_relation_service.Relation{
		PageNum:  util.GetFrontPage(c),
		PageSize: util.GetFrontLimit(c),
		Uid:      uid,
	}
	vo, total, page := relationService.GetUserCollectList()
	response.PageResult(0, vo, "ok", total, page, c)
}
