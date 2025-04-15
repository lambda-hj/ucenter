package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Wechat struct {
		AppID     string
		AppSecret string
		Token     string
		AESKey    string
		NotifyURL string
	}
}
