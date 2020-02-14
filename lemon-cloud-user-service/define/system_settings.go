package define

import (
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components/lccc_define"
	"github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-components/lccc_model"
)

func GetSystemSettings() *lccc_model.SystemSettingsDefine {
	return &lccc_model.SystemSettingsDefine{
		Introduce: "对柠檬云用户中心服务的相关设置",
		SettingGroupList: []*lccc_model.SystemSettingGroupDefine{
			{
				Key:       "system",
				Name:      "系统设置",
				Introduce: "这里包含系统运行所需要的重要相关配置",
				SettingList: []*lccc_model.SystemSettingItemDefine{
					lccc_define.GENERAL_SYSTEM_SETTING_ITEM_DEFINE_DATABASE_TYPE,
					lccc_define.GENERAL_SYSTEM_SETTING_ITEM_DEFINE_DATABASE_URL,
				},
			},
		},
	}
}
