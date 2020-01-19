package client

import "github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-common/service"

func getClient() {
	return service.NewUserLoginServiceClient()
}
