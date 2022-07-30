package main

import (
	"com.go-pro-study/todolist/go-grpc-todolist-study/task/config"
	"com.go-pro-study/todolist/go-grpc-todolist-study/task/discovery"
	"com.go-pro-study/todolist/go-grpc-todolist-study/task/internal/handler"
	"com.go-pro-study/todolist/go-grpc-todolist-study/task/internal/repository"
	"com.go-pro-study/todolist/go-grpc-todolist-study/task/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
)

func main() {
	config.InitConfig()
	respository.InitDB()

	//etcd 地址
	etcdAddress := []string{viper.GetString("etcd.address")}
	//服务注册
	etcdRegister := discovery.NewRegister(etcdAddress, logrus.New())
	grpcAddress := viper.GetString("server.grpcAddress")
	userNode := discovery.Server{
		Name: viper.GetString("server.domain"),
		Addr: grpcAddress,
	}
	server := grpc.NewServer()
	defer server.Stop()

	//绑定服务
	service.RegisterTaskServiceServer(server, handler.NewTaskService())

	listen, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}
	if _, err = etcdRegister.Register(userNode, 10); err != nil {
		panic(err)
	}
	if err = server.Serve(listen); err != nil {
		panic(err)
	}
}
