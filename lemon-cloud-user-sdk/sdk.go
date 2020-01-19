package lemon_cloud_user_sdk

import (
	"fmt"
	"google.golang.org/grpc"
)

var serverConn *grpc.ClientConn

func InitSdk(serverAddress string) {
	con, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		fmt.Println("初始化失败")
	}
	serverConn = con
}

func GetServerConn() *grpc.ClientConn {
	return serverConn
}
