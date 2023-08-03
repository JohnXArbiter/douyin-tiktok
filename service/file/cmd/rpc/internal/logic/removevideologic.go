package logic

import (
	"context"
	"douyin-tiktok/service/file/cmd/rpc/internal/logic/oss"
	"douyin-tiktok/service/file/model"

	"douyin-tiktok/service/file/cmd/rpc/internal/svc"
	"douyin-tiktok/service/file/cmd/rpc/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveVideoLogic {
	return &RemoveVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveVideoLogic) RemoveVideo(in *__.RemoveVideoReq) (*__.CodeResp, error) {
	var (
		ossLogic   = oss.NewOssLogic(l.ctx, l.svcCtx)
		objectName = in.GetObjectName()
	)

	_, err := l.svcCtx.FileVideo.Delete(&model.FileVideo{ObjectName: objectName})
	if err != nil {
		logx.Errorf("[DB ERROR] RemoveVideo 删除视频记录失败 %v\n", err)
		return &__.CodeResp{Code: -1}, err
	}

	ossLogic.Delete(objectName)
	return &__.CodeResp{}, nil
}
