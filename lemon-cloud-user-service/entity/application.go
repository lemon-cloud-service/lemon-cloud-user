package entity

import (
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-model/marshal"
	"github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-service/base"
)

type Application struct {
	base.BaseEntity
	ApplicationKey       string                    `gorm:"size:512"json:"application_key"`
	ApplicationName      string                    `gorm:"size:256" json:"application_name"`
	ApplicationIntroduce string                    `gorm:"size:2048" json:"application_introduce"`
	SecretId             string                    `gorm:"size:512" json:"secret_id"`
	SecretKey            string                    `gorm:"size:2048" json:"secret_key"`
	ExpiredAt            marshal.TimeJsonTimeStamp `json:"expired_at"`
	Capabilities         []ApplicationCapability   `gorm:"ASSOCIATION_FOREIGNKEY:DataKey;FOREIGNKEY:BelongApplicationDataKey" json:"capabilities"`
	Roles                []ApplicationRole         `gorm:"ASSOCIATION_FOREIGNKEY:DataKey;FOREIGNKEY:BelongApplicationDataKey" json:"roles"`
}
