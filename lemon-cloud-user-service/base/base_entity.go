package base

import (
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components/lccc_marshal"
)

type BaseEntity struct {
	DataKey   string                         `gorm:"type:varchar(64);primary_key"`
	CreatedAt lccc_marshal.TimeJsonTimeStamp `json:"created_at"`
	UpdatedAt lccc_marshal.TimeJsonTimeStamp `json:"updated_at"`
	DeletedAt lccc_marshal.TimeJsonTimeStamp `sql:"index" json:"-"`
}
