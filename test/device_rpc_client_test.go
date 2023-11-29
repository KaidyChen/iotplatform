package test

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/zrpc"
	"iot-platform/device/device_client"
	"testing"
)

var deviceRpcClient device_client.Device

func TestSendMessage(t *testing.T) {
	RpcClientConf := zrpc.NewEtcdClientConf([]string{"192.168.2.5:2379"}, "device.rpc", "", "")
	deviceRpcClient = device_client.NewDevice(zrpc.MustNewClient(RpcClientConf))
	reply, err := deviceRpcClient.SendMessage(context.Background(), &device_client.SendMsgRequest{
		ProductKey: "1",
		DeviceKey:  "0002",
		Data:       "hello mqtt",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%v\n", reply)
}
