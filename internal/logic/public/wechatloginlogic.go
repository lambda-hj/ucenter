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

func (l *WechatLoginLogic) WechatLogin(req *types.QrcodeRequest) (resp *types.QrcodeResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
