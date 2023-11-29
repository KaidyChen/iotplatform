package logic

import (
	"context"
	"iot-platform/models"

	"iot-platform/admin/internal/svc"
	"iot-platform/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeviceUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeviceUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeviceUpdateLogic {
	return &DeviceUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeviceUpdateLogic) DeviceUpdate(req *types.DeviceUpdateRequest) (resp *types.DeviceUpdateReply, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.DeviceUpdateReply)
	err = l.svcCtx.DB.Where("identity = ?", req.Identity).Updates(&models.DeviceBasic{
		ProductIdentity: req.ProductIdentity,
		Name:            req.Name,
	}).Error
	if err != nil {
		logx.Error("[ DB ERROR :] ", err)
		resp.StatusCode = "101"
		resp.Msg = err.Error()
		return
	}
	resp.StatusCode = "100"
	resp.Msg = "设备更新成功"
	return
}
