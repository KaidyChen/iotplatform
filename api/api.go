package api

import (
	"encoding/json"
	"errors"
	"iot-platform/define"
	"iot-platform/helper"
)

// emqx服务的api接口，与emqx服务端通信，新建用户和删除用户
func CreateAuthUser(in *CreateAuthUserRequest) error {
	data, _ := json.Marshal(in)
	rep, err := helper.HttpPost(define.EmqxAddr+"/authentication/password_based%3Abuilt_in_database/users", data)
	if err != nil {
		return err
	}
	reply := new(CreateAuthUserReply)
	err = json.Unmarshal(rep, reply)
	if err != nil {
		return errors.New("error client exit")
	}
	return nil
}

func DeleteAuthUser(clientId string) error {
	rep, err := helper.HttpDELETE(define.EmqxAddr+"/authentication/password_based%3Abuilt_in_database/users/"+clientId, []byte{})
	if err != nil {
		return err
	}
	if len(rep) > 0 {
		return errors.New("error client not found")
	}
	return nil
}
