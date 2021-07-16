package handler

import (
	"net/http"

	"zeroService/api/internal/logic"
	"zeroService/api/internal/svc"
	"zeroService/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func userinfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserinfoLogic(r.Context(), ctx)
		resp, err := l.Userinfo(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
