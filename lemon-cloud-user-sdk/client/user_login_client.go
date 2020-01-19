package client

import (
	"github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-common/service"
	"github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-sdk"
)

var defaultUserLoginServiceClient service.UserLoginServiceClient = nil

func GetUserLoginServiceClient() service.UserLoginServiceClient {
	if defaultUserLoginServiceClient == nil {
		defaultUserLoginServiceClient = service.NewUserLoginServiceClient(lemon_cloud_user_sdk.GetServerConn())
	}
	return defaultUserLoginServiceClient
}
