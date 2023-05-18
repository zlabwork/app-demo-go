package network

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
)

const wxEndpoint = "https://api.weixin.qq.com"

type UserSession struct {
	SessionKey string `json:"session_key"`
	OpenId     string `json:"openid"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

type wxHandle struct {
	client *resty.Client
}

func WX() *wxHandle {
	return &wxHandle{
		client: httpClient.SetBaseURL(wxEndpoint),
	}
}

// Code2Session
// @docs https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-login/code2Session.html
// Code2Session {"session_key":"MMRwSObrdVcJue5+INi1uQ==","openid":"o1rzD5Kwpzp-0XJdSPoQ5CJX5pfM"}
func (wx *wxHandle) Code2Session(code string) (*UserSession, error) {

	uri := "/sns/jscode2session"

	var params = map[string]string{
		"appid":      os.Getenv("MINI_APP_ID"),
		"secret":     os.Getenv("MINI_APP_SECRET"),
		"js_code":    code,
		"grant_type": "authorization_code",
	}

	resp, err := wx.client.R().
		SetQueryParams(params).
		Get(uri)

	if err != nil {
		return nil, err
	}

	data := &UserSession{}
	err = json.Unmarshal(resp.Body(), data)
	if err != nil {
		return nil, err
	}
	if data.ErrCode > 0 {
		return nil, fmt.Errorf(data.ErrMsg)
	}

	return data, nil
}
