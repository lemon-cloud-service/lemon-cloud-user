package entity

import "github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-service/base"

type ApplicationCapability struct {
	base.BaseEntity
	CapabilityKey            string `gorm:"size:512" json:"capability_key"`
	CapabilityName           string `gorm:"size:256" json:"capability_name"`
	CapabilityIntroduce      string `gorm:"size:1024" json:"capability_introduce"`
	BelongApplicationDataKey string `gorm:"size:64" json:"belong_application_data_key"`
}
