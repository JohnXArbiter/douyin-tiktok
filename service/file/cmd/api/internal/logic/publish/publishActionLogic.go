package publish

import (
	"context"
	"douyin-tiktok/service/file/cmd/api/internal/logic/oss"
	"douyin-tiktok/service/file/model"
	"errors"
	"mime/multipart"
	"sync"
	"time"

	"douyin-tiktok/service/file/cmd/api/internal/svc"
	"douyin-tiktok/service/file/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishActionLogic {
	return &PublishActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishActionLogic) PublishAction(req *types.PublishActionReq, header *multipart.FileHeader) error {
	var wg sync.WaitGroup

	// TODO 认证
	userIdStr := "123"
	var userId int64 = 123

	ossLogic := oss.NewOssLogic(l.ctx, l.svcCtx)
	url, objectName, err := ossLogic.Upload(header, userIdStr)
	if err != nil {
		return errors.New("视频上传失败！")
	}

	// TODO 调视频（作品）服务rpc更新作品
	var videoId int64
	wg.Add(1)
	go func() {
		defer wg.Done()
	}()

	wg.Wait()
	fv := &model.FileVideo{
		UserId:     userId,
		VideoId:    videoId,
		Objectname: objectName,
		Url:        url,
		UploadAt:   time.Now().Local(),
	}
	_, err = l.svcCtx.FileVideo.Insert(fv)
	if err != nil {
		return errors.New("视频上传失败！")
	}
	return nil
}
