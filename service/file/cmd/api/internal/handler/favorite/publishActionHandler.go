package favorite

import (
	"net/http"

	"douyin-tiktok/service/file/cmd/api/internal/logic/favorite"
	"douyin-tiktok/service/file/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PublishActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := favorite.NewPublishActionLogic(r.Context(), svcCtx)
		err := l.PublishAction()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
