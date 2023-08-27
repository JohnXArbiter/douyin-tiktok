package publish

import (
	"douyin-tiktok/common/middleware"
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/video/cmd/api/internal/logic/publish"
	"douyin-tiktok/service/video/cmd/api/internal/svc"
	"douyin-tiktok/service/video/cmd/api/internal/types"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func PublishActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishActionReq
		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.OkJson(w, utils.GenErrorResp("参数错误！😥"))
			return
		}
		fmt.Println(req.TokenReq)
		loggedUser, err := middleware.JwtAuthenticate(r, req.Token)
		if err != nil {
			httpx.OkJson(w, utils.GenErrorResp(err.Error()))
			return
		}

		// 拿出文件
		_, header, err := r.FormFile("data")
		if err != nil {
			httpx.OkJson(w, utils.GenErrorResp("文件解析错误！😥"))
			return
		}

		l := publish.NewPublishActionLogic(r.Context(), svcCtx)
		err = l.PublishAction(&req, header, loggedUser)
		if err != nil {
			httpx.OkJson(w, utils.GenErrorResp(err.Error()))
		} else {
			httpx.OkJson(w, utils.GenOkResp())
		}
	}
}
