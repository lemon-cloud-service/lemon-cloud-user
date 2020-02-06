package entity

import "github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-service/base"

type UserGroup struct {
	base.BaseEntity
	GroupKey       string            `gorm:"size:512" json:"group_key"`
	GroupName      string            `gorm:"size:256" json:"group_name"`
	GroupIntroduce string            `gorm:"size:1024" json:"group_introduce"`
	Roles          []ApplicationRole `gorm:"many2many:group_roles" json:"roles"`
	Users          []User            `gorm:"many2many:group_users" json:"users"`
}
