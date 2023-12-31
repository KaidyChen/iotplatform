package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mqtt struct {
		Broker   string
		ClientId string
		Password string
	}
}
