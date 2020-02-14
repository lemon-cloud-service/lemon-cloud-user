package entity

import (
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components/lccc_marshal"
	"github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-service/base"
)

type Application struct {
	base.BaseEntity
	ApplicationKey       string                         `gorm:"size:512"json:"application_key"`                                                         // 应用自定义标识
	ApplicationName      string                         `gorm:"size:256" json:"application_name"`                                                       // 应用名称
	ApplicationIntroduce string                         `gorm:"size:2048" json:"application_introduce"`                                                 // 应用介绍
	SecretId             string                         `gorm:"size:512" json:"secret_id"`                                                              // 应用服务间访问的SecretId
	SecretKey            string                         `gorm:"size:2048" json:"secret_key"`                                                            // 应用服务间访问的SecretKey秘钥
	ExpiredAt            lccc_marshal.TimeJsonTimeStamp `json:"expired_at"`                                                                             // 应用过期时间，预留，通常不需要，
	Capabilities         []ApplicationCapability        `gorm:"ASSOCIATION_FOREIGNKEY:DataKey;FOREIGNKEY:BelongApplicationDataKey" json:"capabilities"` // 应用对用户来说所拥有的能力
	Roles                []ApplicationRole              `gorm:"ASSOCIATION_FOREIGNKEY:DataKey;FOREIGNKEY:BelongApplicationDataKey" json:"roles"`        // 应用所提供的角色
}
