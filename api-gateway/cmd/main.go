package main

import (
	"com.go-pro-study/todolist/go-grpc-todolist-study/api-gateway/config"
	"com.go-pro-study/todolist/go-grpc-todolist-study/api-gateway/discovery"
	"com.go-pro-study/todolist/go-grpc-todolist-study/api-gateway/internal/service"
	"com.go-pro-study/todolist/go-grpc-todolist-study/api-gateway/routes"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config.InitConfig()

	//服务发现
	etcdAddress := []string{viper.GetString("etcd.address")}
	etcdRegister := discovery.NewResolver(etcdAddress, logrus.New())
	resolver.Register(etcdRegister)
	go startListen()
	{
		osSignal := make(chan os.Signal, 1)
		signal.Notify(osSignal, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		s := <-osSignal
		fmt.Println("exit!", s)
	}
	fmt.Println("gateway listen on:4000")
}

func startListen() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	userConn, _ := grpc.Dial(viper.GetString("domain.user"), opts...)
	userService := service.NewUserServiceClient(userConn)

	taskConn, _ := grpc.Dial(viper.GetString("domain.task"), opts...)
	taskService := service.NewTaskServiceClient(taskConn)

	//fmt.Println("taskService", taskService)

	ginRouter := routes.NewRouter(userService, taskService)
	server := &http.Server{
		Addr:           viper.GetString("server.port"),
		Handler:        ginRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("绑定失败，端口被占用", err)
	}
}
