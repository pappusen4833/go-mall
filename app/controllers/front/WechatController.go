package front

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/silenceper/wechat/v2/officialaccount/user"
	"go-mall/app/models"
	"go-mall/app/services/wechat_user_service"
	"go-mall/pkg/global"
	"go-mall/pkg/http/response"
	"go-mall/pkg/jwt"
	"go-mall/pkg/redis"
	"go-mall/pkg/util"
	"gorm.io/gorm"
	"io"
	"net/http"
	"strings"
	"time"
)

// WeChat redirect URI is now loaded from configuration

// 公众号服务api
type WechatController struct {
}

// @Title 公众号服务
// @Description 公众号服务
// @Success 200 {object} app.Response
// @router /api/v1/serve [get]
// @Tags Front API
func (e *WechatController) GetAll(c *gin.Context) {
	official := global.OFFICIAL_ACCOUNT
	server := official.GetServer(c.Request, c.Writer)

	server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
		if msg.MsgType == message.MsgTypeEvent {
			global.LOG.Info(msg.Event)
			if msg.Event == message.EventSubscribe {
				//存储用户
				user := official.GetUser()
				userInfo, e := user.GetUserInfo(msg.CommonToken.GetOpenID())
				if e != nil {
					global.LOG.Error(e)
				}
				ip := util.GetClientIP(c)
				userSerive := wechat_user_service.User{UserInfo: userInfo, Ip: ip}
				userSerive.Insert()
			}
		}
		global.LOG.Info(msg.MsgType)
		text := message.NewText(msg.Content)

		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		global.LOG.Error(err)
		return
	}
	//发送回复的消息
	err = server.Send()
	if err != nil {
		global.LOG.Error(err)
		return
	}

}

// LoginUrl 获取登录手机上微信跳转登录的url
func (e *WechatController) LoginUrl(c *gin.Context) {
	key := c.Query("key")
	url := e.getAuthUrl(key)
	response.OkWithData(url, c)
}

func (e *WechatController) LoginResult(c *gin.Context) {
	key := c.Query("key")
	exists := redis.Exists("login:" + key)
	if exists {
		token := redis.GetString("login:" + key)
		token = strings.Replace(token, "\"", "", -1)
		global.LOG.Info("login result data:")
		global.LOG.Info(token)
		data := struct {
			Token string `json:"token"`
		}{Token: token}
		response.OkWithData(data, c)
	} else {
		response.FailWithMessage("", c)
	}
}

func (e *WechatController) getAuthUrl(key string) string {
	appId := global.CONFIG.Wechat.AppID
	redirectUri := global.CONFIG.Wechat.RedirectURI + "?key=" + key
	return "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" + appId + "&redirect_uri=" + redirectUri + "&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect"
}

func (e *WechatController) Verify(c *gin.Context) {
	global.LOG.Info("weixin verify...")
	key := c.Query("key")
	code := c.Query("code")
	global.LOG.Info(key)
	global.LOG.Info("code:," + code)
	userInfo := e.getWexinUserInfo(c, code)
	if userInfo != nil {
		d := time.Now().Add(time.Hour * 24 * 100)
		token, _ := jwt.GenerateAppToken(userInfo, d)
		_ = redis.Set("login:"+key, token, 1000)
		global.LOG.Info("token=" + token)
		frontendURL := global.CONFIG.App.FrontendBaseURL
		c.Redirect(http.StatusMovedPermanently, frontendURL+"/loginResult?flag=1&token="+token)
	} else {
		frontendURL := global.CONFIG.App.FrontendBaseURL
		c.Redirect(http.StatusMovedPermanently, frontendURL+"/loginResult?flag=0")
	}
}

func (e *WechatController) getWexinUserInfo(c *gin.Context, code string) *models.User {
	global.LOG.Info("getWexinUserInfo..")
	appId := global.CONFIG.Wechat.AppID
	appSecret := global.CONFIG.Wechat.AppSecret
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", appId, appSecret, code)
	response, err := http.Get(url)
	if err != nil {
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
	}
	var res map[string]interface{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil
	}
	token := res["access_token"].(string)
	openID := res["openid"].(string)

	urlInfo := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s", token, openID)
	response, err = http.Get(urlInfo)
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	body, err = io.ReadAll(response.Body)
	if err != nil {
		return nil
	}
	userInfo := user.Info{}
	global.LOG.Info(string(body))
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		return nil
	}
	global.LOG.Info(userInfo)
	ip := util.GetClientIP(c)

	var u models.User
	err = global.DB.
		Model(&models.User{}).
		Where("username = ?", userInfo.OpenID).First(&u).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		global.LOG.Info("user not found, create")
		// 未找到用户 需要新建
		userService := wechat_user_service.User{UserInfo: &userInfo, Ip: ip}
		err = userService.Insert()
		if err != nil {
			// 新建用户失败
			global.LOG.Info("create err")
			global.LOG.Info(err)
		} else {
			// 新建用户成功
			global.LOG.Info("no error")
			global.EVENT_BUS.Publish("user:create", userInfo.OpenID)
		}
	}
	_ = global.DB.
		Model(&models.User{}).
		Where("username = ?", userInfo.OpenID).First(&u).Error
	return &u
}
