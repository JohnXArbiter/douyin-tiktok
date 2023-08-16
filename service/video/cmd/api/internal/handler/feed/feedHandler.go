package feed

import (
	"douyin-tiktok/common/middleware"
	"douyin-tiktok/common/utils"
	"fmt"
	"net/http"

	"douyin-tiktok/service/video/cmd/api/internal/logic/feed"
	"douyin-tiktok/service/video/cmd/api/internal/svc"
	"douyin-tiktok/service/video/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FeedHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FeedReq
		if err := httpx.ParseForm(r, &req); err != nil {
			fmt.Println(err)
			httpx.OkJson(w, utils.GenErrorResp("参数错误！😥"))
			return
		}

		loggedUser, err := middleware.JwtAuthenticate(r, req.Token)
		if err != nil {
			httpx.OkJson(w, utils.GenErrorResp(err.Error()))
			return
		}

		l := feed.NewFeedLogic(r.Context(), svcCtx)
		if resp, err := l.Feed(&req, loggedUser); err != nil {
			httpx.OkJson(w, utils.GenErrorResp(err.Error()))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
