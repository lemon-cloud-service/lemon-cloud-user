package entity

import "github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-service/base"

type UserGroup struct {
	base.BaseEntity
	GroupKey           string            `gorm:"size:512" json:"group_key"`            // 用户组自定义标识
	GroupName          string            `gorm:"size:256" json:"group_name"`           // 用户组名称
	GroupIntroduce     string            `gorm:"size:1024" json:"group_introduce"`     // 用户组介绍
	Roles              []ApplicationRole `gorm:"many2many:group_roles" json:"roles"`   // 用户组所拥有权限
	Users              []User            `gorm:"many2many:group_users" json:"users"`   // 组内用户，多对多关系
	ParentGroupDataKey string            `gorm:"size:64" json:"parent_group_data_key"` // 父级用户组的DataKey，如果是根层，那么留空字符串
}
