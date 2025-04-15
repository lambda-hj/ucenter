package svc

import (
	"ucenter/internal/config"
	"ucenter/pkg/wechat"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount"
)

type ServiceContext struct {
	Config config.Config
	Wechat *officialAccount.OfficialAccount
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Wechat: wechat.MustNewWeChatClient(c.Wechat.AppID, c.Wechat.AppSecret, c.Wechat.Token, c.Wechat.AESKey, c.Wechat.NotifyURL),
	}
}
