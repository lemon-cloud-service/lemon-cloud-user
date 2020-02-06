package define

import (
	"fmt"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-define/lccd_strings"
)

var SYSTEM_INFO_NAME string = "Lemon Cloud User Center"
var SYSTEM_INFO_VERSION string = "1.0.0"
var SYSTEM_INFO_SPLIT_LINE string = "====================================================================="

// 打印系统的基础信息，包含LemonCloud字符画和系统名称及版本
func PrintSystemInfo() {
	fmt.Print(lccd_strings.LEMON_CLOUD_ASCII_IMAGE)
	fmt.Println(SYSTEM_INFO_SPLIT_LINE)
	fmt.Printf("%v [ver: %v] is starting...\n", SYSTEM_INFO_NAME, SYSTEM_INFO_VERSION)
	fmt.Println(SYSTEM_INFO_SPLIT_LINE)
}
