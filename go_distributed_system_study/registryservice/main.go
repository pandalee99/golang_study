package main

import (
	"context"
	"fmt"
	"go_distributed_system_study/registry"
	"log"
	"net/http"
)

func main() {
	registry.SetupRegistryService()
	//将之前的处理逻辑注册进去

	http.Handle("/services", &registry.RegistryService{})

	//接下来的逻辑一样，需要有取消功能，当然其实你在ide中能直接打断，但在大型服务中，每秒都要运行。还是需要自定义取消功能的
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//一样是定义服务的地址
	var srv http.Server
	srv.Addr = registry.ServerPort

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Println("注册中心 的服务开始。按任意键停止. \n")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("结束服务注册")
}
