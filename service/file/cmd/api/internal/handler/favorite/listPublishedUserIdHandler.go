package favorite

import (
	"net/http"

	"douyin-tiktok/service/file/cmd/api/internal/logic/favorite"
	"douyin-tiktok/service/file/cmd/api/internal/svc"
	"douyin-tiktok/service/file/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListPublishedUserIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := favorite.NewListPublishedUserIdLogic(r.Context(), svcCtx)
		err := l.ListPublishedUserId(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
