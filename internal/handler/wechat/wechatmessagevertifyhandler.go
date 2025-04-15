package wechat

import (
	"net/http"

	"ucenter/internal/logic/wechat"
	"ucenter/internal/svc"

	"github.com/ArtisanCloud/PowerLibs/v3/http/helper"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func WechatMessageVertifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := wechat.NewWechatMessageVertifyLogic(r.Context(), svcCtx)
		rsp, err := l.WechatMessageVertify(r.Context(), r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			helper.HttpResponseSend(rsp, w)
		}
	}
}
