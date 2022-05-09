package service

import (
	"context"
	"fmt"
	"go_distributed_system_study/registry"
	"log"
	"net/http"
)

//公共的函数，用于启动服务
func Start(ctx context.Context, host, port string, reg registry.Registration,
	registerHandlersFunc func()) (context.Context, error) {
	//将传入的函数运行
	registerHandlersFunc()

	//对服务进行基本的定义，完善服务，并将信息返回给主函数
	ctx = startService(ctx, reg.ServiceName, host, port)

	//启动web服务之后注册：
	err := registry.RegisterService(reg)
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

func startService(ctx context.Context, name registry.ServiceName, host string, port string) context.Context {
	//使得ctx具有取消的功能
	ctx, cancel := context.WithCancel(ctx)

	//定义服务地址
	var server http.Server
	server.Addr = ":" + port

	//
	go func() {
		//启动服务，一旦发生可错误，就取消
		log.Println(server.ListenAndServe())
		//调用取消服务的服务
		err := registry.ShutdownService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			log.Println(err)
		}
		cancel()
	}()

	go func() {
		fmt.Printf("%v 服务开始。按任意键停止. \n", name)
		var s string
		fmt.Scanln(&s)

		//调用取消服务的服务
		err := registry.ShutdownService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			log.Println(err)
		}

		server.Shutdown(ctx)
		cancel()
	}()

	return ctx
}
