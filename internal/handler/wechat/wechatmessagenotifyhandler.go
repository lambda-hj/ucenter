package wechat

import (
	"net/http"

	"ucenter/internal/logic/wechat"
	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 微信消息回调消息与验证
func WechatMessageNotifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := wechat.NewWechatMessageNotifyLogic(r.Context(), svcCtx)
		rsp, err := l.WechatMessageNotify(r.Context(), r)
		l.Infof("rsp: %+v", rsp)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
