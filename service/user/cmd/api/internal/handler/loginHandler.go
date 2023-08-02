package handler

import (
	"douyin-tiktok/common/utils"
	"net/http"

	"douyin-tiktok/service/user/cmd/api/internal/logic"
	"douyin-tiktok/service/user/cmd/api/internal/svc"
	"douyin-tiktok/service/user/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func loginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJson(w, utils.GenErrorResp("参数错误！😥"))
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.OkJson(w, utils.GenErrorResp(err.Error()))
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
