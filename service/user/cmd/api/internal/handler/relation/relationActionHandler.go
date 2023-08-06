package relation

import (
	"douyin-tiktok/common/middleware"
	"douyin-tiktok/common/utils"
	"fmt"
	"net/http"

	"douyin-tiktok/service/user/cmd/api/internal/logic/relation"
	"douyin-tiktok/service/user/cmd/api/internal/svc"
	"douyin-tiktok/service/user/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RelationActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RelationActionReq
		if err := httpx.ParseJsonBody(r, &req); err != nil {
			fmt.Println(req, err)
			httpx.OkJson(w, utils.GenErrorResp("参数错误！😥"))
			return
		}

		loggedUser, err := middleware.JwtAuthenticate(r, req.Token)
		if err != nil {
			httpx.OkJson(w, utils.GenErrorResp(err.Error()))
			return
		}

		l := relation.NewRelationActionLogic(r.Context(), svcCtx)
		err = l.RelationAction(&req, loggedUser)
		if err != nil {
			httpx.OkJson(w, utils.GenErrorResp(err.Error()))
		} else {
			httpx.OkJson(w, utils.GenOkResp())
		}
	}
}
