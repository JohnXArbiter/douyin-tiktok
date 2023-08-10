package message

import (
	"douyin-tiktok/common/middleware"
	"douyin-tiktok/common/utils"
	"net/http"

	"douyin-tiktok/service/user/cmd/api/internal/logic/message"
	"douyin-tiktok/service/user/cmd/api/internal/svc"
	"douyin-tiktok/service/user/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatReq
		if err := httpx.ParseForm(r, &req); err != nil {
			httpx.OkJson(w, utils.GenErrorResp("参数错误！😥"))
			return
		}

		loggedUser, err := middleware.JwtAuthenticate(r, req.Token)
		if err != nil {
			httpx.OkJson(w, utils.GenErrorResp(err.Error()))
			return
		}

		l := message.NewChatLogic(r.Context(), svcCtx)
		if resp, err := l.Chat(&req, loggedUser); err != nil {
			httpx.OkJson(w, utils.GenErrorResp(err.Error()))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
