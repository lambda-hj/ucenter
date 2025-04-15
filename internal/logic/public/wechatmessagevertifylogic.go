package public

import (
	"context"
	"net/http"

	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type WechatMessageVertifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWechatMessageVertifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WechatMessageVertifyLogic {
	return &WechatMessageVertifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WechatMessageVertifyLogic) WechatMessageVertify(ctx context.Context, r *http.Request) error {
	_, err := l.svcCtx.Wechat.Server.VerifyURL(r)
	if err != nil {
		return err
	}
	return nil
}
