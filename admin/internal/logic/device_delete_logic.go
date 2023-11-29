package logic

import (
	"context"
	"gorm.io/gorm"
	"iot-platform/admin/internal/svc"
	"iot-platform/admin/internal/types"
	"iot-platform/api"
	"iot-platform/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceDeleteLogic {
	return &DeviceDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceDeleteLogic) DeviceDelete(req *types.DeviceDeleteRequest) (resp *types.DeviceDeleteReply, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.DeviceDeleteReply)
	deviceBasic := new(models.DeviceBasic)
	var count int64
	//先查找数据库中是否存在目标设备
	err = l.svcCtx.DB.Model(new(models.DeviceBasic)).Select("key").Where("identity = ?", req.Identity).
		Find(deviceBasic).Count(&count).Error
	if err != nil {
		logx.Error("[ DB ERROR]: ", err)
		resp.StatusCode = "101"
		resp.Msg = err.Error()
		return
	}
	if count == 0 {
		resp.StatusCode = "101"
		resp.Msg = "设备不存在"
		return
	}
	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		//1. 数据库中删除设备
		err = tx.Where("identity = ?", req.Identity).Unscoped().Delete(new(models.DeviceBasic)).Error
		if err != nil {
			logx.Error("[DB ERROR]:", err)
			return err
		}
		//2. EMQX 中同步删除认证设备
		err = api.DeleteAuthUser(deviceBasic.Key)
		if err != nil {
			logx.Error("[DB ERROR]:", err)
			return err
		}
		return nil
	})
	if err != nil {
		resp.StatusCode = "101"
		resp.Msg = err.Error()
	}
	resp.StatusCode = "100"
	resp.Msg = "设备删除成功"
	return
}
