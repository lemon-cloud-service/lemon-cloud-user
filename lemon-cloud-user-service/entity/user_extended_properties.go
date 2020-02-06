package entity

import "github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-service/base"

type UserExtendedPropertiesDefine struct {
	base.BaseEntity
	PropertyKey       string `gorm:"size:512" json:"property_key"`
	PropertyName      string `gorm:"size:256" json:"property_name"`
	PropertyIntroduce string `gorm:"size:1024" json:"property_introduce"`
}

type UserExtendedPropertiesStorage struct {
	base.BaseEntity
	UserDataKey           string `gorm:"size:64" json:"user_data_key"`
	PropertyDefineDataKey string `gorm:"size:64" json:"properties_define_data_key"`
	PropertyValue         string `gorm:"size:2048" json:"property_value"`
}
