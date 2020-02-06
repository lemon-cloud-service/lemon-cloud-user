package base

import (
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-model/marshal"
)

type BaseEntity struct {
	DataKey   string                    `gorm:"type:varchar(64);primary_key"`
	CreatedAt marshal.TimeJsonTimeStamp `json:"created_at"`
	UpdatedAt marshal.TimeJsonTimeStamp `json:"updated_at"`
	DeletedAt marshal.TimeJsonTimeStamp `sql:"index" json:"-"`
}
