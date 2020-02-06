package entity

import (
	"github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-service/base"
)

type User struct {
	base.BaseEntity
	Number           string      `gorm:"size:128" json:"number"`
	Salt             string      `gorm:"size:64" json:"salt"`
	Password         string      `gorm:"size:512" json:"password"`
	Name             string      `gorm:"size:512" json:"name"`
	NickName         string      `gorm:"size:512" json:"nick_name"`
	Birthday         uint64      `json:"birthday"`
	Phone            string      `gorm:"size:64" json:"phone"`
	Email            string      `gorm:"size:256" json:"email"`
	BelongUserGroups []UserGroup `gorm:"many2many:group_users" json:"belong_user_groups"`
}
