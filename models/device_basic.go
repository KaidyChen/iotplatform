package models

import (
	"gorm.io/gorm"
	"time"
)

type DeviceBasic struct {
	gorm.Model
	Identity        string `gorm:"column:identity; type:varchar(50);" json:"identity"`
	ProductIdentity string `gorm:"column:product_identity; type:varchar(50);" json:"product_identity"`
	Name            string `gorm:"column:name; type:varchar(50);" json:"name"`
	Key             string `gorm:"column:key; type:varchar(50);" json:"key"`
	Secret          string `gorm:"column:secret; type:varchar(50);" json:"secret"`
	LastOnlineTime  string `gorm:"column:last_online_time; type:varchar(50);" json:"last_online_time"`
}

func (table DeviceBasic) TableName() string {
	return "device_basic"
}

// GetDeviceList 获取设备列表
func GetDeviceList(name string) *gorm.DB {
	tx := DB.Model(new(DeviceBasic)).Select("device_basic.identity, device_basic.name," +
		"device_basic.key, device_basic.secret, pb.name product_name, device_basic.last_online_time").
		Joins("LEFT JOIN product_basic pb ON pb.identity = device_basic.product_identity")
	if name != "" {
		tx.Where("device_basic.name LIKE ?", "%"+name+"%")
	}
	return tx
}

// UpdateDeviceOnlineTime 更新设备上线时间
func UpdateDeviceOnlineTime(productkey, deviceKey string) error {
	var productIdentity string
	err := DB.Model(new(ProductBasic)).Select("identity").Where("`key` = ?", productkey).Scan(&productIdentity).Error
	if err != nil {
		return err
	}
	err = DB.Model(new(DeviceBasic)).Where("`key` = ? AND product_identity = ?", deviceKey, productkey).Update("last_online_time",
		time.Now().Format("2006-01-02 15:04:05")).Error
	return err
}
