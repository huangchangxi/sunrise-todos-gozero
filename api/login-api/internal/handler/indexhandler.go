package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"login-api/internal/logic"
	"login-api/internal/svc"
)

func IndexHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewIndexLogic(r.Context(), ctx)
		resp, err := l.Index()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
