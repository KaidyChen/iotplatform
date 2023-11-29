package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"iot-platform/device/device_client"
	"iot-platform/open/internal/config"
	"iot-platform/user/rpc/user_client"
)

type ServiceContext struct {
	Config    config.Config
	RpcDevice device_client.Device
	RpcUser   user_client.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		RpcDevice: device_client.NewDevice(zrpc.MustNewClient(c.RpcDevice)),
		RpcUser:   user_client.NewUser(zrpc.MustNewClient(c.RpcUser)),
	}
}
