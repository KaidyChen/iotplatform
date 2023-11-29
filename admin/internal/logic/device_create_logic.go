package logic

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"iot-platform/api"
	"iot-platform/helper"
	"iot-platform/models"

	"iot-platform/admin/internal/svc"
	"iot-platform/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceCreateLogic {
	return &DeviceCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceCreateLogic) DeviceCreate(req *types.DeviceCreateRequest) (resp *types.DeviceCreateReply, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.DeviceCreateReply)
	//先查找所属产品是否存在
	var count int64
	l.svcCtx.DB.Where("identity = ?", req.ProductIdentity).First(&models.ProductBasic{}).Count(&count)
	if count == 0 {
		resp.StatusCode = "101"
		resp.Msg = "所属产品类型不存在, 请检查产品类型"
		return
	}
	l.svcCtx.DB.Where("name = ?", req.Name).Find(&models.DeviceBasic{}).Count(&count)
	fmt.Println("count:", count)
	if count != 0 {
		resp.StatusCode = "102"
		resp.Msg = "设备已存在，请勿重复创建"
		return
	}
	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		//1. 数据库中新增设备
		deviceBasic := &models.DeviceBasic{
			Identity:        uuid.New().String(),
			ProductIdentity: req.ProductIdentity,
			Name:            req.Name,
			Key:             uuid.New().String(),
			Secret:          uuid.New().String(),
		}
		err = tx.Create(deviceBasic).Error
		if err != nil {
			logx.Error("[DB ERROR] : ", err)
			return err
		}
		//2. EMQX 中新增认证设备
		err = api.CreateAuthUser(&api.CreateAuthUserRequest{
			UserId:   deviceBasic.Key,
			Password: helper.Md5(deviceBasic.Key + deviceBasic.Secret),
		})
		if err != nil {
			logx.Error("[CreateAuthUse ERROR] : ", err)
			return err
		}
		//提交事务
		return nil
	})
	if err != nil {
		resp.StatusCode = "101"
		resp.Msg = err.Error()
		return
	}
	resp.StatusCode = "100"
	resp.Msg = "设备创建成功"
	return
}
