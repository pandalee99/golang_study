package main

import (
	"context"
	"go_distributed_system_study/grades"
	"go_distributed_system_study/log"
	"go_distributed_system_study/registry"
	"go_distributed_system_study/service"

	"fmt"
	stlog "log"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	r := registry.Registration{
		ServiceName: registry.GradingService,
		ServiceURL:  serviceAddress,
		//添加两个信息
		RequiredServices: []registry.ServiceName{registry.LogService},
		ServiceUpdateURL: serviceAddress + "/services",
		HeartbeatURL:     serviceAddress + "/heartbeat",
	}
	ctx, err := service.Start(context.Background(),
		host,
		port,
		r,
		grades.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}

	//在服务启动之后使用log服务
	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("发现日志服务: %s\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}

	<-ctx.Done()
	fmt.Println("grading service 服务停止了")
}
