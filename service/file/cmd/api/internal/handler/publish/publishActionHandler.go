package publish

import (
	"errors"
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"douyin-tiktok/service/file/cmd/api/internal/logic/publish"
	"douyin-tiktok/service/file/cmd/api/internal/svc"
	"douyin-tiktok/service/file/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PublishActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishActionReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New("参数错误！😥"))
			return
		}

		// 拿出文件
		_, header, err := r.FormFile("data")
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New("文件解析错误！😥"))
			return
		}

		l := publish.NewPublishActionLogic(r.Context(), svcCtx)
		err = l.PublishAction(&req, header)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, nil)
		}
	}
}
