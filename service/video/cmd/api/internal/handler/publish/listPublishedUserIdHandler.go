package publish

import (
	"errors"
	xhttp "github.com/zeromicro/x/http"
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
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New("参数错误！😥"))
			return
		}

		l := publish.NewListPublishedUserIdLogic(r.Context(), svcCtx)
		resp, err := l.ListPublishedUserId(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, nil)
		}
	}
}
