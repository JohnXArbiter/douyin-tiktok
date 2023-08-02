package publish

import (
	"douyin-tiktok/common/utils"
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
			httpx.OkJson(w, utils.GenErrorResp("参数错误！😥"))
			return
		}

		// 拿出文件
		_, header, err := r.FormFile("data")
		if err != nil {
			httpx.OkJson(w, utils.GenErrorResp("文件解析错误！😥"))
			return
		}

		l := publish.NewPublishActionLogic(r.Context(), svcCtx)
		err = l.PublishAction(&req, header)
		if err != nil {
			httpx.OkJson(w, utils.GenErrorResp(err.Error()))
		} else {
			httpx.OkJson(w, utils.GenOkResp())
		}
	}
}
