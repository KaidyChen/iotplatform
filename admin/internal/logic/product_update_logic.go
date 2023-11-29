package logic

import (
	"context"
	"iot-platform/models"

	"iot-platform/admin/internal/svc"
	"iot-platform/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductUpdateLogic {
	return &ProductUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductUpdateLogic) ProductUpdate(req *types.ProductUpdateRequest) (resp *types.ProductUpdateReply, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.ProductUpdateReply)
	err = l.svcCtx.DB.Where("identity = ?", req.Identity).Updates(&models.ProductBasic{
		Name: req.Name,
		Desc: req.Desc,
	}).Error
	if err != nil {
		logx.Error("[DB ERROR]:", err)
		resp.StatusCode = "101"
		resp.Msg = err.Error()
	}
	resp.StatusCode = "100"
	resp.Msg = "设备更新成功"
	return
}
