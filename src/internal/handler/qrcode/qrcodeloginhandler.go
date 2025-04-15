package qrcode

import (
	"net/http"

	"ucenter/src/internal/logic/qrcode"
	"ucenter/src/internal/svc"
	"ucenter/src/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取微信登录二维码
func QrcodeLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QrcodeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := qrcode.NewQrcodeLoginLogic(r.Context(), svcCtx)
		resp, err := l.QrcodeLogin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
