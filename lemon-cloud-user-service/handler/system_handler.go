package handler

import (
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components/lccc_micro_service"
	"github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-service/define"
	"github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-service/manager"
	log "github.com/sirupsen/logrus"
	"os"
)

func SystemStart() {
	var err error

	// print system info
	define.PrintSystemInfo()

	// read config file
	log.Info("Start reading configuration files...")
	err = manager.ConfigManagerInstance().Init()
	if err != nil {
		log.Error("System start failed. Error reading configuration file: ", err.Error())
		os.Exit(1)
	}
	log.Info("Configuration file read completed")

	// repair config data

	// init registry
	log.Info("Start configuring the registry...")
	err = lccc_micro_service.SystemServiceInstance().RegisterNewService(&lccc_micro_service.ServiceRegisterConfig{
		ServiceGeneralConfig: manager.ConfigManagerInstance().GeneralConfig(),
		ServiceInfo:          define.GetServiceInfo(),
	})
	if err != nil {
		log.Error("System start failed. Error configuring registry: ", err.Error())
		os.Exit(1)
	}
	log.Info("Registry config completed")
}
