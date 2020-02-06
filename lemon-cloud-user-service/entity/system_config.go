package entity

import "github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-service/base"

type SystemConfig struct {
	base.BaseEntity
	ConfigKey       string `gorm:"size:512" json:"config_key"`
	ConfigName      string `gorm:"size:256" json:"config_name"`
	ConfigIntroduce string `gorm:"size:1024" json:"config_introduce"`
	ConfigValue     string `gorm:"size:10240" json:"config_value"`
}
