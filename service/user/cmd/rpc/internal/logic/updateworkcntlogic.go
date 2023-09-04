package logic

import (
	"context"
	"douyin-tiktok/service/user/cmd/rpc/internal/svc"
	"douyin-tiktok/service/user/cmd/rpc/types"
	"douyin-tiktok/service/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateWorkCntLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateWorkCntLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateWorkCntLogic {
	return &UpdateWorkCntLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateWorkCntLogic) UpdateWorkCnt(in *__.UpdateWorkCntReq) (*__.CodeResp, error) {
	session := l.svcCtx.UserInfo
	if in.IsAdd == 1 {
		session.SetExpr("work_count", "work_count + 1")
	} else {
		session.SetExpr("work_count", "work_count - 1").Where("work_count > 0")
	}

	if _, err := session.Update(&model.UserInfo{Id: in.UserId}); err != nil {
		logx.Errorf("[DB ERROR] UpdateWorkCnt 更新作品输失败 %v\n", err)
		return &__.CodeResp{Code: -1}, err
	}

	return &__.CodeResp{}, nil
}
