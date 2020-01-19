package main

import (
	"google.golang.org/grpc"
	"lemonit.cn/cloud/user/service/service"
	"lemonit.cn/cloud/user/service/service_impl"
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
	service.RegisterEnrollServiceServer(s, service_impl.EnrollServiceImpl{})
	s.Serve(listen)
}
