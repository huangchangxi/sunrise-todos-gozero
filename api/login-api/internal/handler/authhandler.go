package handler

import (
	"net/http"

	"login-api/internal/logic"
	"login-api/internal/svc"
	"login-api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func AuthHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AuthReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAuthLogic(r.Context(), ctx)
		resp, err := l.Auth(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
