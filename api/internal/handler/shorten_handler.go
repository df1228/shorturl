package handler

import (
	"net/http"

	"github.com/dfang/shorturl/api/internal/logic"
	"github.com/dfang/shorturl/api/internal/svc"
	"github.com/dfang/shorturl/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShortenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShortenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewShortenLogic(r.Context(), svcCtx)
		resp, err := l.Shorten(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
