package handler

import (
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
	err = manager.EtcdManagerInstance().ClientInit(
		manager.ConfigManagerInstance().GeneralConfig().Registry.Endpoints,
		manager.ConfigManagerInstance().GeneralConfig().Registry.Username,
		manager.ConfigManagerInstance().GeneralConfig().Registry.Password)
	if err != nil {
		log.Error("System start failed. Error configuring registry: ", err.Error())
		os.Exit(1)
	}
	log.Info("Registry config completed")
}
