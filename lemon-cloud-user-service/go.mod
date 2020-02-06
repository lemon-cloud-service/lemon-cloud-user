module github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-service

go 1.13

require (
	github.com/jinzhu/gorm v1.9.12
	github.com/lemon-cloud-service/lemon-cloud-common v0.0.0-20200206134200-f189d76cf056
	github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-model v0.0.0-20200206095924-351975e38e75
	github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils v0.0.0-20200206055610-d888d2092c7c
	github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-common v0.0.0-20200119034614-e04a7f7a61d1
	github.com/sirupsen/logrus v1.4.2
	google.golang.org/grpc v1.26.0
)

replace github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-common => ../lemon-cloud-user-common
