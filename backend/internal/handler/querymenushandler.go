package handler

import (
	"net/http"

	"WePublicMenu/internal/logic"
	"WePublicMenu/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryMenusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewQueryMenusLogic(r.Context(), svcCtx)
		resp, err := l.QueryMenus()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
