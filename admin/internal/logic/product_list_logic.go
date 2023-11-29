package logic

import (
	"context"
	"iot-platform/helper"
	"iot-platform/models"

	"iot-platform/admin/internal/svc"
	"iot-platform/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductListLogic {
	return &ProductListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductListLogic) ProductList(req *types.ProductListRequest) (resp *types.ProductListReply, err error) {
	// todo: add your logic here and delete this line
	productList := make([]*types.ProductListBasic, 0)
	req.Size = helper.If(req.Size == 0, 20, req.Size).(int)
	req.Page = helper.If(req.Page == 0, 0, (req.Page-1)*req.Size).(int)
	resp = new(types.ProductListReply)
	var count int64
	err = models.ProductList(req.Name).Count(&count).Offset(req.Page).Limit(req.Size).Find(&productList).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err)
		return
	}
	for _, v := range productList {
		v.CreatedAt = helper.RFC3339ToNormalTime(v.CreatedAt)
	}
	resp.List = productList
	resp.Count = count
	return
}
