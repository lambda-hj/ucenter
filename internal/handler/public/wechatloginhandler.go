package public

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ucenter/internal/logic/public"
	"ucenter/internal/svc"
	"ucenter/internal/types"
)

// 获取微信登录二维码
func WechatLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QrcodeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := public.NewWechatLoginLogic(r.Context(), svcCtx)
		resp, err := l.WechatLogin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
