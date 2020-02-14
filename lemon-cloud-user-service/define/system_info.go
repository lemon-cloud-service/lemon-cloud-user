package define

import (
	"fmt"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components/lccc_define"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components/lccc_model"
)

const SYSTEM_INFO_NAME string = "Lemon Cloud User Center"
const SYSTEM_INFO_SERVICE_TAG string = "lemon_cloud_user"
const SYSTEM_INFO_SERVICE_INTRODUCE string = "Lemon Cloud User Center Service"
const SYSTEM_INFO_VERSION string = "1.0.0"
const SYSTEM_INFO_SPLIT_LINE string = "====================================================================="

func GetServiceInfo() *lccc_model.ServiceInfo {
	return &lccc_model.ServiceInfo{
		ServiceTag:       SYSTEM_INFO_SERVICE_TAG,
		ServiceName:      SYSTEM_INFO_NAME,
		ServiceIntroduce: SYSTEM_INFO_SERVICE_INTRODUCE,
	}
}

// 打印系统的基础信息，包含LemonCloud字符画和系统名称及版本
func PrintSystemInfo() {
	fmt.Print(lccc_define.LEMON_CLOUD_ASCII_IMAGE)
	fmt.Print("\n")
	fmt.Println(SYSTEM_INFO_SPLIT_LINE)
	fmt.Printf("Welcome to %v [ver: %v]\n", SYSTEM_INFO_NAME, SYSTEM_INFO_VERSION)
	fmt.Println(SYSTEM_INFO_SPLIT_LINE)
}
