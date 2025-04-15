package svc

import (
	"context"
	"ucenter/src/internal/config"
	"ucenter/src/pkg/wechat"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount"
	"github.com/zeromicro/go-zero/core/logc"
)

type ServiceContext struct {
	Config config.Config
	wechat *officialAccount.OfficialAccount
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		wechat: wechat.MustNewWeChatClient(c.Wechat.AppID, c.Wechat.AppSecret, c.Wechat.Token, c.Wechat.AESKey, c.Wechat.NotifyURL),
	}
}

func (ctx *ServiceContext) GetWechatQRcodeURL(c context.Context, machineID string) (string, error) {
	rsp, err := ctx.wechat.QRCode.Temporary(c, machineID, 60*30)
	if err != nil {
		return "", err
	}
	logc.Infof(c, "wechat qrcode ticket: %s", rsp.Ticket)
	return rsp.Url, nil
}
