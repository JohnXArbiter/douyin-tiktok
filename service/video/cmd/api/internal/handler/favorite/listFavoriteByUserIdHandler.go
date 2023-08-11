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

func ListFavoriteByUserIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, utils.GenErrorResp("请求参数错误"))
			return
		}

		if _, err := middleware.JwtAuthenticate(r, req.Token); err != nil {
			httpx.OkJson(w, utils.GenErrorResp(err.Error()))
			return
		}

		l := favorite.NewListFavoriteByUserIdLogic(r.Context(), svcCtx)
		if resp, err := l.ListFavoriteByUserId(&req); err != nil {
			httpx.OkJson(w, utils.GenErrorResp(err.Error()))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
