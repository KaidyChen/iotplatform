package logic

import (
	"context"
	"errors"
	"iot-platform/models"

	"iot-platform/admin/internal/svc"
	"iot-platform/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDeleteLogic {
	return &ProductDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductDeleteLogic) ProductDelete(req *types.ProductDeleteRequest) (resp *types.ProductDeleteReply, err error) {
	// todo: add your logic here and delete this line
	var count int64
	resp = new(types.ProductDeleteReply)
	err = l.svcCtx.DB.Model(new(models.DeviceBasic)).Where("product_identity = ?", req.Identity).Count(&count).Error
	if err != nil {
		logx.Error("[DB ERROR]:", err)
		resp.StatusCode = "102"
		resp.Msg = err.Error()
		return
	}
	if count > 0 {
		err = errors.New("已关联设备，请先删除关联设备")
		resp.StatusCode = "103"
		resp.Msg = err.Error()
	}
	err = l.svcCtx.DB.Where("identity = ?", req.Identity).Unscoped().Delete(new(models.ProductBasic)).Error
	if err != nil {
		logx.Error("[DB ERROR]:", err)
		resp.StatusCode = "102"
		resp.Msg = err.Error()
	}
	resp.StatusCode = "100"
	resp.Msg = "产品删除成功"
	return
}
