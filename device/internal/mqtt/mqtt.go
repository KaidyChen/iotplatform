package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/zeromicro/go-zero/core/logx"
	"iot-platform/models"
	"log"
	"strings"
)

var topic = "/sys/#"
var MC mqtt.Client

func publishHandler(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("MESSAGE: %s\n", message.Payload())
	fmt.Printf("TOPIC: %s\n", message.Topic())
	//to do 当收到设备心跳消息时更新设备最新在线时间
	topicArray := strings.Split(strings.TrimPrefix(message.Topic(), "/"), "/")
	if len(topicArray) >= 4 {
		if topicArray[3] == "ping" {
			err := models.UpdateDeviceOnlineTime(topicArray[1], topicArray[2])
			if err != nil {
				logx.Error("[DB ERROR]:", err)
			}
		}
	}
}

func NewMqttServer(mattBroker, clientId, password string) {
	opt := mqtt.NewClientOptions().AddBroker(mattBroker).SetClientID(clientId).SetUsername("get").SetPassword(password)

	//设置回调
	opt.SetDefaultPublishHandler(publishHandler)

	MC = mqtt.NewClient(opt)

	//连接
	if token := MC.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	//订阅主题
	if token := MC.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	defer func() {
		//取消订阅
		if token := MC.Unsubscribe(topic); token.Wait() && token.Error() != nil {
			log.Println("[ERROR]: ", token.Error())
		}
		//关闭连接
		MC.Disconnect(250)
	}()

	//开始等待
	select {}
}
