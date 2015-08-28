package weixin

import (
	"encoding/json"
	"fmt"
	"github.com/huyan/beego/httplib"
)

// get user info
func (wx *Weixin) GetUserInfo(openId string, v interface{}) error {
	reply, err := sendGetRequest(fmt.Sprintf(weixinHost+"/user/info?openid=%s&access_token=", openId), wx.tokenChan)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(reply, v); err != nil {
		return err
	}
	return nil
}

type WxAccessTokenResp struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string
	Scope        string
}

func GetUserInfoOAuth1(appid, appsecret, code string) (token, openid string, err error) {
	accessTokenUrl := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	accessTokenUrl = fmt.Sprintf(accessTokenUrl, appid, appsecret, code)

	wxResp := &WxAccessTokenResp{}
	err = httplib.Get(accessTokenUrl).ToJson(wxResp)
	if err != nil {
		return "", "", fmt.Errorf("oauth 获取认证access_token错误：%s,%s", wxResp.Openid, err)
	}
	return wxResp.AccessToken, wxResp.Openid, nil
}

func GetUserInfoOAuth2(appid, appsecret, token, Openid string, v interface{}) error {
	usrInfoUrl := "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN"
	usrInfoUrl = fmt.Sprintf(usrInfoUrl, token, Openid)
	err := httplib.Get(usrInfoUrl).ToJson(v)
	if err != nil {
		return fmt.Errorf("oauth 获取授权用户信息错误：%s,%s", Openid, err)
	}
	return nil
}
