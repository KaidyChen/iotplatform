package logic

import (
	"context"
	"errors"
	"iot-platform/helper"
	"iot-platform/user/rpc/internal/svc"
	"iot-platform/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLogic {
	return &AuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthLogic) Auth(in *user.UserAuthRequest) (*user.UserAuthReply, error) {
	// todo: add your logic here and delete this line
	if in.Token == "" {
		return nil, errors.New("token不能为空")
	}
	userClaim, err := helper.AnalyzeToken(in.Token)
	if err != nil {
		return nil, err
	}
	return &user.UserAuthReply{
		Identity: userClaim.Identity,
		Id:       uint64(userClaim.Id),
		Extend: map[string]string{
			"name": userClaim.Name,
		},
	}, nil
}
