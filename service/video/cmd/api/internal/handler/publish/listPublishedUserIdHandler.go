package publish

import (
	"douyin-tiktok/common/middleware"
	"douyin-tiktok/common/utils"
	"net/http"

	"douyin-tiktok/service/video/cmd/api/internal/logic/publish"
	"douyin-tiktok/service/video/cmd/api/internal/svc"
	"douyin-tiktok/service/video/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListPublishedUserIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, utils.GenErrorResp("参数错误！😥"))
			return
		}

		err := middleware.JwtAuthenticate(r, req.Token)
		if err != nil {
			httpx.OkJson(w, utils.GenErrorResp(err.Error()))
			return
		}

		l := publish.NewListPublishedUserIdLogic(r.Context(), svcCtx)
		resp, err := l.ListPublishedUserId(&req)
		if err != nil {
			httpx.OkJson(w, utils.GenErrorResp(err.Error()))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
