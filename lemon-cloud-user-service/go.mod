module github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-service

go 1.13

require (
	github.com/coreos/etcd v3.3.18+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-00010101000000-000000000000 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-define v0.0.0-00010101000000-000000000000
	github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-model v0.0.0-20200206095924-351975e38e75
	github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils v0.0.0-20200206055610-d888d2092c7c
	github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-common v0.0.0-20200119034614-e04a7f7a61d1
	github.com/sirupsen/logrus v1.4.2
	go.etcd.io/etcd v3.3.18+incompatible
	go.uber.org/zap v1.13.0 // indirect
)

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0

replace github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-common => ../lemon-cloud-user-common

replace github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-define => ../../lemon-cloud-common/lemon-cloud-common-define

replace github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-utils => ../../lemon-cloud-common/lemon-cloud-common-utils

replace github.com/lemon-cloud-service/lemon-cloud-common/lemon-cloud-common-model => ../../lemon-cloud-common/lemon-cloud-common-model
