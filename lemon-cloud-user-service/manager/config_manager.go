package manager

import (
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-define/lccd_strings"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-model/lccm_config"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils/lccu_config"
	"sync"
)

type ConfigManager struct {
	generalConfig *lccm_config.GeneralConfig
}

var configManagerInstance *ConfigManager
var configManagerOnce sync.Once

func ConfigManagerInstance() *ConfigManager {
	configManagerOnce.Do(func() {
		configManagerInstance = &ConfigManager{}
	})
	return configManagerInstance
}

func (cm *ConfigManager) Init() error {
	if !cm.InitializationStatus() {
		cm.generalConfig = &lccm_config.GeneralConfig{}
		err := lccu_config.LoadYamlConfigFile(lccd_strings.FILE_PATH_GENERAL_CONFIG_FILE, cm.generalConfig)
		if err != nil {
			return err
		}
	}
	return nil
}

func (cm *ConfigManager) GeneralConfig() *lccm_config.GeneralConfig {
	return cm.generalConfig
}

func (cm *ConfigManager) InitializationStatus() bool {
	return cm.generalConfig != nil
}
