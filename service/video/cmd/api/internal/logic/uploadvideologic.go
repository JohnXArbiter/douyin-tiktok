package logic

//
//import (
//	"bytes"
//	"context"
//	"douyin-tiktok/service/file/cmd/rpc/internal/logic/oss"
//	"douyin-tiktok/service/file/cmd/rpc/internal/svc"
//	"douyin-tiktok/service/file/cmd/rpc/types"
//	"douyin-tiktok/service/file/model"
//	"fmt"
//	ffmpeg "github.com/u2takey/ffmpeg-go"
//	"github.com/zeromicro/go-zero/core/logx"
//	"io"
//	"os"
//	"strconv"
//	"strings"
//	"time"
//)
//
//type UploadVideoLogic struct {
//	ctx    context.Context
//	svcCtx *svc.ServiceContext
//	logx.Logger
//}
//
//func NewUploadVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadVideoLogic {
//	return &UploadVideoLogic{
//		ctx:    ctx,
//		svcCtx: svcCtx,
//		Logger: logx.WithContext(ctx),
//	}
//}
//
//func (l *UploadVideoLogic) UploadVideo(in *__.UploadVideoReq) (*__.UploadVideoResp, error) {
//	var (
//		resp      = &__.UploadVideoResp{Code: -1}
//		userId    = in.UserId
//		userIdStr = strconv.FormatInt(userId, 10)
//		videoFile = bytes.NewReader(in.GetVideoFile())
//		videoId   = in.VideoId
//		fileName  = in.VideoName
//	)
//
//	session := l.svcCtx.Xorm.NewSession()
//	defer session.Close()
//
//	if err := session.Begin(); err != nil {
//		logx.Errorf("[DB ERROR] UploadVideo 开启事务失败 %v\n", err)
//		return resp, err
//	}
//
//	ossLogic := oss.NewOssLogic(l.ctx, l.svcCtx)
//	videoUrl, videoObjectName, err := ossLogic.UploadRaw(videoFile, fileName, userIdStr)
//	if err != nil {
//		return resp, err
//	}
//
//	coverFile, err := l.getFrameByFfmpeg(videoUrl)
//	coverUrl, coverObjectName, err := ossLogic.UploadRaw(coverFile, fileName[:strings.LastIndex(fileName, ".")]+".png", userIdStr)
//	if err != nil {
//		return resp, err
//	}
//
//	now := time.Now().Unix()
//	fileVideo := &model.FileVideo{
//		UserId:     userId,
//		VideoId:    videoId,
//		ObjectName: videoObjectName,
//		Url:        videoUrl,
//		UploadAt:   now,
//	}
//	if _, err = l.svcCtx.FileVideo.Insert(fileVideo); err != nil {
//		go ossLogic.Delete(videoObjectName)
//		logx.Errorf("[DB ERROR] UploadVideo 插入视频文件记录失败 %v\n", err)
//		return resp, err
//	}
//
//	fileCover := &model.FileCover{
//		UserId:     userId,
//		VideoId:    videoId,
//		ObjectName: coverObjectName,
//		Url:        coverUrl,
//		UploadAt:   now,
//	}
//	if _, err = l.svcCtx.FileCover.Insert(fileCover); err != nil {
//		go ossLogic.Delete(coverObjectName)
//		logx.Errorf("[DB ERROR] UploadVideo 插入封面文件记录失败 %v\n", err)
//		return resp, err
//	}
//
//	if err = session.Commit(); err != nil {
//		logx.Errorf("[DB ERROR] UploadVideo 提交事务失败 %v\n", err)
//		go ossLogic.Delete(coverObjectName)
//		go ossLogic.Delete(videoObjectName)
//	}
//
//	resp.Code = 0
//	resp.PlayUrl = videoUrl
//	resp.CoverUrl = coverUrl
//	return resp, nil
//}
//
//func (l *UploadVideoLogic) getFrameByFfmpeg(videoUrl string) (io.Reader, error) {
//
//	buf := bytes.NewBuffer(nil)
//	err := ffmpeg.Input(videoUrl).
//		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
//		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
//		WithOutput(buf, os.Stdout).
//		Run()
//
//	if err != nil {
//		logx.Errorf("[FFMPEG ERROR] 生成视频封面图 %v\n", err)
//		return nil, err
//	}
//	return buf, nil
//	//img, err := imaging.Decode(buf)
//	//if err != nil {
//	//	log.Fatal("222生成缩略图失败：", err)
//	//	return err
//	//}
//	//
//	//err = imaging.Save(img, "fuck.png")
//	//if err != nil {
//	//	log.Fatal("333生成缩略图失败：", err)
//	//	return err
//	//}
//	//return nil
//}
