package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"login-api/internal/logic"
	"login-api/internal/svc"
)

func VerifyCodeHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewVerifyCodeLogic(r.Context(), ctx)
		resp, err := l.VerifyCode()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
