package logic

import (
	"context"
	"encoding/json"
	"errors"
	"iot-platform/helper"
	"iot-platform/models"
	"iot-platform/user/rpc/internal/svc"
	"iot-platform/user/rpc/types/user"
	"sort"

	"github.com/zeromicro/go-zero/core/logx"
)

type OpenAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOpenAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OpenAuthLogic {
	return &OpenAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OpenAuthLogic) OpenAuth(in *user.OpenAuthRequest) (*user.OpenAuthReply, error) {
	// todo: add your logic here and delete this line
	data := make(map[string]interface{})
	err := json.Unmarshal(in.Data, &data)
	if err != nil {
		logx.Error("[Unmarshal ERROR] : ", err.Error())
		return nil, err
	}
	userbasic := new(models.UserBasic)
	err = l.svcCtx.Db.Model(new(models.UserBasic)).Select("app_secret").Where("app_key = ?", data["app_key"]).
		First(userbasic).Error
	if err != nil {
		logx.Error("[DB ERROR] : ", err.Error())
		return nil, err
	}
	//判断用户信息是否存在
	if userbasic.AppSecret == "" {
		logx.Error("Invalid AppKey")
		return nil, err
	}
	dataArry := make([]string, 0)
	for k, _ := range data {
		dataArry = append(dataArry, k)
	}
	sort.Strings(dataArry)
	var sign string
	for _, v := range dataArry {
		if v != "sign" {
			sign += data[v].(string)
		}
	}
	if helper.Md5(sign) != data["sign"].(string) {
		return nil, errors.New("签名不正确")
	}
	return &user.OpenAuthReply{}, nil
}
