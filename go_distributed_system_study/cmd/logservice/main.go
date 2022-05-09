package main

import (
	"context"
	"fmt"
	"go_distributed_system_study/log"
	"go_distributed_system_study/registry"
	"go_distributed_system_study/service"
	stlog "log"
)

func main() {
	//定义日志地址
	log.Run("./distribute.log")
	//定义具体参数，其实通常应该由配置文件中定义
	host, port := "localhost", "4000"

	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)
	r := registry.Registration{
		ServiceName: registry.LogService,
		ServiceURL:  serviceAddress,
		//添加两个信息
		RequiredServices: make([]registry.ServiceName, 0),
		ServiceUpdateURL: serviceAddress + "/services",
		HeartbeatURL:     serviceAddress + "/heartbeat",
	}

	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		log.RegisterHandler,
	)
	//如果有错误，就先执行标准库的log打印出结果
	if err != nil {
		stlog.Fatalln(err)
	}
	//接需要等待ctx的信号
	//如果在启动服务器的时候出现了错误
	//或者在按下任意键停止后，就会发送信号
	<-ctx.Done()
	//接受到信号后，就会继续
	fmt.Println("停止服务")
}
