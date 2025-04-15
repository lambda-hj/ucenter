package public

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"ucenter/internal/logic/public"
	"ucenter/internal/svc"
)

func WechatMessageVertifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := public.NewWechatMessageVertifyLogic(r.Context(), svcCtx)
		err := l.WechatMessageVertify(r.Context(), r)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
