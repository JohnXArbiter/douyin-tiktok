package svc

import (
	"douyin-tiktok/common/utils"
	"douyin-tiktok/service/file/cmd/rpc/fileservice"
	"douyin-tiktok/service/user/cmd/rpc/userservice"
	"douyin-tiktok/service/video/cmd/api/internal/config"
	"github.com/yitter/idgenerator-go/idgen"
	"github.com/zeromicro/go-zero/zrpc"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config

	UserRpc userservice.UserService
	FileRpc fileservice.FileService

	Xorm          *xorm.Engine
	VideoInfo     *xorm.Session
	VideoFavorite *xorm.Session
	VideoComment  *xorm.Session
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine := utils.InitXorm("mysql", c.Mysql)

	options := idgen.NewIdGeneratorOptions(c.Idgen.WorkerId)
	idgen.SetIdGenerator(options)

	return &ServiceContext{
		Config:        c,
		Xorm:          engine,
		VideoInfo:     engine.Table("video_info"),
		VideoFavorite: engine.Table("video_favorite"),
		VideoComment:  engine.Table("video_comment"),
		UserRpc:       userservice.NewUserService(zrpc.MustNewClient(c.UserRpc)),
		FileRpc:       fileservice.NewFileService(zrpc.MustNewClient(c.FileRpc)),
		//Redis:        utils.InitRedis(c.Redis),
	}
}
