package qrcode

import (
	"context"

	"ucenter/src/internal/svc"
	"ucenter/src/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QrcodeLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取微信登录二维码
func NewQrcodeLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QrcodeLoginLogic {
	return &QrcodeLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QrcodeLoginLogic) QrcodeLogin(req *types.QrcodeRequest) (resp *types.QrcodeResponse, err error) {
	notifyURL, err := l.svcCtx.GetWechatQRcodeURL(l.ctx, req.Machine_id)
	if err != nil {
		return nil, err
	}
	return &types.QrcodeResponse{
		Message: "success",
		Url:     notifyURL,
	}, nil
}
