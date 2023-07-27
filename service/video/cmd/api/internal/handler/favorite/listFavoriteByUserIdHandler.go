package favorite

import (
	"net/http"

	"douyin-tiktok/service/video/cmd/api/internal/logic/favorite"
	"douyin-tiktok/service/video/cmd/api/internal/svc"
	"douyin-tiktok/service/video/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListFavoriteByUserIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := favorite.NewListFavoriteByUserIdLogic(r.Context(), svcCtx)
		err := l.ListFavoriteByUserId(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
