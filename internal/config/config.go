package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Wechat struct {
		AppID     string `json:",env=WECHAT_APP_ID"`
		AppSecret string `json:",env=WECHAT_APP_SECRET"`
		Token     string `json:",env=WECHAT_TOKEN"`
		AESKey    string `json:",env=WECHAT_AES_KEY"`
		NotifyURL string
	}
}
