package wechat

import (
	"context"
	"strings"

	"io"
	"net/http"

	"ucenter/internal/svc"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/messages"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/server/handlers/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type WechatMessageNotifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 微信消息回调消息与验证
func NewWechatMessageNotifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WechatMessageNotifyLogic {
	return &WechatMessageNotifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WechatMessageNotifyLogic) WechatMessageNotify(ctx context.Context, r *http.Request) (string, error) {
	rsp, err := l.svcCtx.Wechat.Server.Notify(r, func(event contract.EventInterface) interface{} {
		l.Debugf("event: %+v", event)
		switch event.GetMsgType() {
		case models.CALLBACK_EVENT_SCAN:
			msg := models.EventScanCodePush{}
			err := event.ReadMessage(&msg)
			if err != nil {
				return "error"
			}
			open_id := msg.FromUserName
			envet_key := msg.EventKey
			machine_id := strings.Split(envet_key, "_")[1]
			l.Infof("open_id: %s, machine_id: %s", open_id, machine_id)
			return messages.NewText("ok")
		case models.CALLBACK_EVENT_SUBSCRIBE:
			msg := models.EventSubscribe{}
			err := event.ReadMessage(&msg)
			if err != nil {
				return "error"
			}
			return messages.NewText("ok")
		case models.CALLBACK_EVENT_UNSUBSCRIBE:
			msg := models.EventUnSubscribe{}
			err := event.ReadMessage(&msg)
			if err != nil {
				return "error"
			}
			return messages.NewText("ok")
		}
		return kernel.SUCCESS_EMPTY_RESPONSE

	})

	if err != nil {
		return "", err
	}
	text, _ := io.ReadAll(rsp.Body)
	return string(text), nil
}
