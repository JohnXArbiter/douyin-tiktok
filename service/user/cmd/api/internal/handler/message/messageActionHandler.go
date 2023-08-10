package message

import (
	"douyin-tiktok/common/middleware"
	"douyin-tiktok/common/utils"
	"net/http"
	"strings"

	"douyin-tiktok/service/user/cmd/api/internal/logic/message"
	"douyin-tiktok/service/user/cmd/api/internal/svc"
	"douyin-tiktok/service/user/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func MessageActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MessageAction
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, utils.GenErrorResp("参数错误！😥"))
			return
		} else if strings.TrimSpace(req.Content) == "" {
			httpx.OkJson(w, utils.GenErrorResp("输入内容不能为空！"))
			return
		}

		loggedUser, err := middleware.JwtAuthenticate(r, req.Token)
		if err != nil {
			httpx.OkJson(w, utils.GenErrorResp(err.Error()))
			return
		}

		l := message.NewMessageActionLogic(r.Context(), svcCtx)
		if err = l.MessageAction(&req, loggedUser); err != nil {
			httpx.OkJson(w, utils.GenErrorResp(err.Error()))
		} else {
			httpx.OkJson(w, utils.GenOkResp())
		}
	}
}
