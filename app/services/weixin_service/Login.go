package weixin_service

import (
	"encoding/json"
	"fmt"
	"go-mall/pkg/global"
	"io/ioutil"
	"net/http"
)

func getAuthUrl(key string) string {
	appId := global.CONFIG.Wechat.AppID
	redirectUri := global.CONFIG.Wechat.RedirectURI + "?key=" + key
	return "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" + appId + "&redirect_uri=" + redirectUri + "&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect"
}

func Verify(key string, code string) {
	// 获取用户信息
	//res := GetUserInfo(code)
	//if res != nil {
	//Cache.Store("redis").Set("login:"+key, res, 300)
	//token := res["token"]
	//frontendURL := global.CONFIG.App.FrontendBaseURL
	//http.Redirect(w, r, frontendURL+"/loginResult?flag=1&token="+token, http.StatusSeeOther)
	//} else {
	//frontendURL := global.CONFIG.App.FrontendBaseURL
	//http.Redirect(w, r, frontendURL+"/loginResult?flag=0", http.StatusSeeOther)
	//}
}

func GetUserInfo(code string) string {
	// Get WeChat credentials from config
	appId := global.CONFIG.Wechat.AppID
	appSecret := global.CONFIG.Wechat.AppSecret
	url := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + appId + "&secret=" + appSecret + "&code=" + code + "&grant_type=authorization_code"
	//global.LOG.Info("weixin user," + url)
	//json对象变成数组
	res := make(map[string]interface{})
	response, err := http.Get(url)
	if err != nil {
		//global.LOG.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		//global.LOG.Fatal(err)
	}
	json.Unmarshal(body, &res)
	// 获取用户信息

	errcode := res["errcode"].(int)
	if errcode != 0 {
		fmt.Print("error")
	}
	token := res["access_token"].(string)
	openId := res["openid"].(string)
	urlInfo := "https://api.weixin.qq.com/sns/userinfo?access_token=" + token + "&openid=" + openId
	response, err = http.Get(urlInfo)
	if err != nil {
		//global.LOG.Fatal(err)
	}
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		//global.LOG.Fatal(err)
	}

	//user := models.User{
	//	Username:  u.UserInfo.OpenID,
	//	Nickname:  u.UserInfo.Nickname,
	//	Password:  util.HashAndSalt([]byte("123456")),
	//	RealName:  u.UserInfo.Nickname,
	//	Avatar:    u.UserInfo.Headimgurl,
	//	AddIp:     u.Ip,
	//	LastIp:    u.Ip,
	//	UserType:  userEnum.WECHAT,
	//	WxProfile: datatypes.JSON(result),
	//}
	//wxUserInfo := make(map[string]interface{})
	//json.Unmarshal(body, &wxUserInfo)
	//global.LOG.info("wxUserInfo," + string(body))
	//// 判断用户是否存在, 否则新增
	//res = wxUser.where("openid", openId).find()
	//var member *Member
	//if res != nil {
	//	// 更新用户
	//	wxUserInfo["id"] = res.id
	//	wxUser.update(wxUserInfo)
	//	// 登陆用户
	//	member = Member.where("id", res.user_id).find()
	//	memberId := member.id
	//	session.Set("member", member) //浏览器关闭断开失效
	//} else {
	//	// 生成账号并登陆
	//	memberId := memberService.createMember(1)
	//	if memberId != nil {
	//		wxUserInfo["user_id"] = memberId
	//	} else {
	//		return false
	//	}
	//	// 保存数据
	//	result := this.insertData(wxUserInfo)
	//	if !result {
	//		return nil
	//	}
	//}
	//if memberId == nil {
	//	return nil
	//}
	//// 生成jwt
	//jUtil := new(JwtUtil)
	//token := jUtil.createJwt(memberId)
	//return map[string]interface{}{
	//	"token":  token,
	//	"member": member,
	//	"wxuser": wxUserInfo,
	//}
	return ""
}
