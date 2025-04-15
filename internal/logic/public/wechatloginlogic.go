package public

import (
	"context"

	"ucenter/internal/svc"
	"ucenter/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WechatLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取微信登录二维码
func NewWechatLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WechatLoginLogic {
	return &WechatLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WechatLoginLogic) WechatLogin(ctx context.Context, req *types.QrcodeRequest) (resp *types.QrcodeResponse, err error) {
	rsp, err := l.svcCtx.Wechat.QRCode.Temporary(ctx, req.Machine_id, 60*30)
	if err != nil {
		return nil, err
	}
	l.Infof("rsp: %+v", rsp)
	return &types.QrcodeResponse{
		Message: "ok",
		Url:     rsp.Url,
	}, nil
}
