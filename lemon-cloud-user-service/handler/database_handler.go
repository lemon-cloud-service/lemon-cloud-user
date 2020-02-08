package handler

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils/lccu_strings"
	"strings"
)

func InitDatabaseConfig() {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "LCU_" + strings.ToUpper(lccu_strings.CamelToSnake(defaultTableName))
	}
}

func AutoProcessingTables() {

}
