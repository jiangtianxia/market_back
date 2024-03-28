package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	// 微信小程序appid
	wxAppid string

	// 微信小程序secret
	wxSecret string
)

const (
	// 获取微信openid URL格式
	getOpenidURLFormat = `https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code`
)

func InitWXAPI(appid, secret string) {
	wxAppid = appid
	wxSecret = secret
}

type WXGetOpenidResp struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// GetOpenid 获取微信openid
func GetOpenid(code string) (string, error) {
	// 构建url
	url := fmt.Sprintf(getOpenidURLFormat, wxAppid, wxSecret, code)

	// 发起http请求
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read body failed, err: %v", err)
	}

	// 解析结果并返回
	var res WXGetOpenidResp
	err = json.Unmarshal(body, &res)
	if err != nil {
		return "", fmt.Errorf("unmarshal body failed, body: %v, err: %v", body, err)
	}
	if res.ErrCode != 0 {
		return "", fmt.Errorf("get openid failed, errcode: %d, errmsg: %s", res.ErrCode, res.ErrMsg)
	}
	return res.Openid, nil
}
