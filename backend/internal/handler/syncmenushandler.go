package handler

import (
	"net/http"

	"WePublicMenu/internal/logic"
	"WePublicMenu/internal/svc"
	"WePublicMenu/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SyncMenusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SyncMenusRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSyncMenusLogic(r.Context(), svcCtx)
		resp, err := l.SyncMenus(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
