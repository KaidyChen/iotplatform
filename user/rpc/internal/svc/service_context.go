package svc

import (
	"gorm.io/gorm"
	"iot-platform/models"
	"iot-platform/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.NewDB()
	return &ServiceContext{
		Config: c,
		Db:     models.DB,
	}
}
