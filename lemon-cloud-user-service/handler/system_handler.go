package handler

import (
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components/lccc_micro_service"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils/lccu_log"
	"github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-service/define"
	"github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-service/manager"
	"os"
)

func SystemStart() {
	var err error

	// print system info
	define.PrintSystemInfo()

	// read config file
	lccu_log.Info("Start reading configuration files...")
	err = manager.ConfigManagerInstance().Init()
	if err != nil {
		lccu_log.Error("System start failed. Error reading configuration file: ", err.Error())
		os.Exit(1)
	}
	lccu_log.Info("Configuration file read completed")

	// repair config data

	// init registry
	lccu_log.Info("Start configuring the registry...")
	err = lccc_micro_service.CoreServiceSingletonInstance().RegisterServiceInstance(&lccc_micro_service.ServiceRegisterConfig{
		ServiceGeneralConfig:   manager.ConfigManagerInstance().GeneralConfig(),
		ServiceBaseInfo:        define.GetServiceBaseInfo(),
		ServiceApplicationInfo: define.GetServiceApplicationInfo(),
	}, define.GetSystemSettings())
	if err != nil {
		lccu_log.Error("System start failed. Error configuring registry: ", err.Error())
		os.Exit(1)
	}
	lccu_log.Info("Registry config completed")
}
