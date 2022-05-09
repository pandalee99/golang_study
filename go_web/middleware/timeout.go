package middleware

import (
	"context"
	"net/http"
	"time"
)

type TimeoutMiddleware struct {
	Next http.Handler
}

func (tm TimeoutMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if tm.Next == nil {
		tm.Next = http.DefaultServeMux
	} //处理和中间件的一般方式一样

	ctx := request.Context()                         //获取上下文
	ctx, _ = context.WithTimeout(ctx, 3*time.Second) //修改context的超时判断
	request.WithContext(ctx)                         //用我们自定义的context去代替
	ch := make(chan struct{})                        //意图在于，如果我们请求能够在3秒内完成的话，这个chan就会收到一个信号
	go func() {                                      //使用goroutine
		tm.Next.ServeHTTP(writer, request) //执行完这个方法后，发送一个信号
		ch <- struct{}{}                   //发送信号
	}()
	select { //一个竞争的状态
	case <-ch: //正常处理完，得到信号，返回
		return
	case <-ctx.Done(): //否则返回错误
		writer.WriteHeader(http.StatusRequestTimeout)
	}
	ctx.Done()
}
