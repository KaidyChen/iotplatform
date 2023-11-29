package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"iot-platform/admin/internal/config"
	"iot-platform/models"
	"iot-platform/user/rpc/types/user"
	"iot-platform/user/rpc/user_client"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	RpcUser user.UserClient
	Auth    *user.UserAuthReply
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.NewDB()
	return &ServiceContext{
		Config:  c,
		DB:      models.DB,
		RpcUser: user_client.NewUser(zrpc.MustNewClient(c.RpcUser)),
	}
}
