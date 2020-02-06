package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils/lccu_strings"
	"strings"
)

func InitDatabaseConfig() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "LCU_" + strings.ToUpper(lccu_strings.CamelToSnake(defaultTableName))
	}
}
