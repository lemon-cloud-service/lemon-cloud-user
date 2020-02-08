package define

import (
	"fmt"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-define/lccd_strings"
)

const SYSTEM_INFO_NAME string = "Lemon Cloud User Center"
const SYSTEM_INFO_VERSION string = "1.0.0"
const SYSTEM_INFO_SPLIT_LINE string = "====================================================================="

// 打印系统的基础信息，包含LemonCloud字符画和系统名称及版本
func PrintSystemInfo() {
	fmt.Print(lccd_strings.LEMON_CLOUD_ASCII_IMAGE)
	fmt.Print("\n")
	fmt.Println(SYSTEM_INFO_SPLIT_LINE)
	fmt.Printf("Welcome to %v [ver: %v]\n", SYSTEM_INFO_NAME, SYSTEM_INFO_VERSION)
	fmt.Println(SYSTEM_INFO_SPLIT_LINE)
}
