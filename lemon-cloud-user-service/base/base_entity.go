package base

import (
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-model/lccm_marshal"
)

type BaseEntity struct {
	DataKey   string                         `gorm:"type:varchar(64);primary_key"`
	CreatedAt lccm_marshal.TimeJsonTimeStamp `json:"created_at"`
	UpdatedAt lccm_marshal.TimeJsonTimeStamp `json:"updated_at"`
	DeletedAt lccm_marshal.TimeJsonTimeStamp `sql:"index" json:"-"`
}
