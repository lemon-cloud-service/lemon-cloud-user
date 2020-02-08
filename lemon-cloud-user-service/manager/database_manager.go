package manager

import (
	"errors"
	"github.com/jinzhu/gorm"
	"sync"
)

type DatabaseManager struct {
	dbInstance *gorm.DB
}

var databaseManagerInstance *DatabaseManager
var databaseManagerOnce sync.Once

func DatabaseManagerInstance() *DatabaseManager {
	databaseManagerOnce.Do(func() {
		databaseManagerInstance = &DatabaseManager{}
	})
	return databaseManagerInstance
}

func (dbm *DatabaseManager) DbInit(dbType, dbUrl string) error {
	if !dbm.DbInitializationStatus() {
		instance, err := gorm.Open(dbType, dbUrl)
		if err != nil {
			return err
		}
		dbm.dbInstance = instance
		defer dbm.dbInstance.Close()
		return nil
	} else {
		return errors.New("the database has been initialized, if you need to reconfigure the database, please restart the system")
	}
}

func (dbm *DatabaseManager) DbInstance() *gorm.DB {
	return dbm.dbInstance
}

func (dbm *DatabaseManager) DbInitializationStatus() bool {
	return dbm.dbInstance != nil
}
