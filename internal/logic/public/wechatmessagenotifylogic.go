package public

import (
	"context"
	"io"
	"net/http"

	"ucenter/internal/svc"

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
			return messages.NewText("ok")
		default:
			return messages.NewText("ok")
		}

	})

	if err != nil {
		return "", err
	}
	text, _ := io.ReadAll(rsp.Body)
	return string(text), nil
}
