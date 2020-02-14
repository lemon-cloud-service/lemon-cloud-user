package entity

import "github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-service/base"

type UserExtendedPropertiesDefine struct {
	base.BaseEntity
	PropertyKey       string `gorm:"size:512" json:"property_key"`        // 扩展属性自定义标识
	PropertyName      string `gorm:"size:256" json:"property_name"`       // 扩展属性名称
	PropertyIntroduce string `gorm:"size:1024" json:"property_introduce"` // 扩展属性介绍
}

type UserExtendedPropertiesStorage struct {
	base.BaseEntity
	UserDataKey           string `gorm:"size:64" json:"user_data_key"`              // 扩展属性所绑定的用户
	PropertyDefineDataKey string `gorm:"size:64" json:"properties_define_data_key"` // 关联的扩展属性DataKey
	PropertyValue         string `gorm:"size:2048" json:"property_value"`           // 指定用户的扩展属性值
}
