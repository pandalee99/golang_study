package middleware

import "net/http"

//中间件
type AuthMiddleware struct {
	Next http.Handler
}

func (am *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if am.Next == nil {
		am.Next = http.DefaultServeMux
	} //如果什么都没有，执行默认页面
	auth := request.Header.Get("Authorization")
	if auth != "" { //说明存在身份，则执行逻辑
		am.Next.ServeHTTP(writer, request)
	} else { //否则报异常
		writer.WriteHeader(http.StatusUnauthorized) //401未授权
	}
}
