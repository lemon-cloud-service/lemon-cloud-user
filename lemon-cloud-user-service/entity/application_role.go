package entity

import "github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-service/base"

type ApplicationRole struct {
	base.BaseEntity
	RoleKey                  string                  `gorm:"size:512" json:"role_key"`
	RoleName                 string                  `gorm:"size:256" json:"role_name"`
	RoleIntroduce            string                  `gorm:"size:1024" json:"role_introduce"`
	BelongApplicationDataKey string                  `gorm:"size:64" json:"belong_application_data_key"`
	Capabilities             []ApplicationCapability `gorm:"many2many:role_capabilities" json:"capabilities"`
	BelongUserGroups         []UserGroup             `gorm:"many2many:group_roles" json:"belong_user_groups"`
}
