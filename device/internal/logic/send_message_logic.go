package logic

import (
	"context"
	"errors"
	"iot-platform/device/internal/mqtt"

	"iot-platform/device/internal/svc"
	"iot-platform/device/types/device"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMessageLogic) SendMessage(in *device.SendMsgRequest) (*device.SendMsgReply, error) {
	// todo: add your logic here and delete this line
	if in.ProductKey == "" || in.DeviceKey == "" || in.Data == "" {
		return nil, errors.New("参数异常")
	}
	topic := "/sys/" + in.ProductKey + "/" + in.DeviceKey + "/receive"
	if token := mqtt.MC.Publish(topic, 0, false, in.Data); token.Wait() && token.Error() != nil {
		logx.Error("[PUBLISH ERROR]:", token.Error())
	}
	return &device.SendMsgReply{}, nil
}
