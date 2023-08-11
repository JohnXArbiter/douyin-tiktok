package favorite

import (
	"douyin-tiktok/common/middleware"
	"douyin-tiktok/common/utils"
	"net/http"

	"douyin-tiktok/service/video/cmd/api/internal/logic/favorite"
	"douyin-tiktok/service/video/cmd/api/internal/svc"
	"douyin-tiktok/service/video/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FavoriteActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FavoriteActionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, utils.GenErrorResp("参数错误！😥"))
			return
		}

		loggedUser, err := middleware.JwtAuthenticate(r, req.Token)
		if err != nil {
			httpx.OkJson(w, utils.GenErrorResp(err.Error()))
			return
		}

		l := favorite.NewFavoriteActionLogic(r.Context(), svcCtx)
		if err := l.FavoriteAction(&req, loggedUser); err != nil {
			httpx.OkJson(w, utils.GenErrorResp(err.Error()))
		} else {
			httpx.OkJson(w, utils.GenOkResp())
		}
	}
}
