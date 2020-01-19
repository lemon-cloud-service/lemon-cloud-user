package main

import (
	"github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-common/service"
	"github.com/lemon-cloud-service/lemon-cloud-user/lemon-cloud-user-service/service_impl"
	"google.golang.org/grpc"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":33385")
	if err != nil {
		println("failed to listen: %v", err)
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册HelloService
	service.RegisterUserLoginServiceServer(s, &service_impl.LoginServiceImpl{})
	s.Serve(listen)
}
