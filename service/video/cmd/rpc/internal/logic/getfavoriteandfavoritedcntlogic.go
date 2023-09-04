package logic

import (
	"context"
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/video/model"
	"strconv"

	"douyin-tiktok/service/video/cmd/rpc/internal/svc"
	"douyin-tiktok/service/video/cmd/rpc/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFavoriteAndFavoritedCntLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFavoriteAndFavoritedCntLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFavoriteAndFavoritedCntLogic {
	return &GetFavoriteAndFavoritedCntLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFavoriteAndFavoritedCntLogic) GetFavoriteAndFavoritedCnt(in *__.GetFavoriteAndFavoritedCntReq) (*__.GetFavoriteAndFavoritedCntResp, error) {
	var (
		userId    = in.UserId
		userIdStr = strconv.FormatInt(userId, 10)
	)

	totalFavorited, err := l.svcCtx.VideoInfo.Where("user_id = ?", 1111).SumInt(&model.VideoInfo{}, "favorite_count")
	if err != nil {
		logx.Errorf("[DB ERROR] GetFavoriteAndFavoritedCnt 查询视频点赞总数失败 %v\n", err)
		return &__.GetFavoriteAndFavoritedCntResp{Code: -1}, err
	}

	favoriteCommonLogic := NewFavoriteCommonLogic(l.ctx, l.svcCtx)
	ids, err := favoriteCommonLogic.LoadIdsAndStore(utils.VideoFavorite+userIdStr, userId)
	if err != nil {
		logx.Errorf("[MONGO ERROR] GetFavoriteAndFavoritedCnt 查询视频点赞总数失败 %v\n", err)
		return &__.GetFavoriteAndFavoritedCntResp{Code: -1}, err
	}

	resp := &__.GetFavoriteAndFavoritedCntResp{
		TotalFavorited: totalFavorited,
		FavoriteCount:  int64(len(ids)),
	}
	return resp, nil
}
