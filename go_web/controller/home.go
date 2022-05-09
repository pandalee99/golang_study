package controller

import "net/http"

func registerHomeRoutes() {
	http.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		if pusher, ok := writer.(http.Pusher); ok { //先做类型断言，看看存不存在push
			pusher.Push("/css/app.css", &http.PushOptions{ //如果为真，则说明支持server push
				Header: http.Header{"Content-Type": []string{"text/css"}}, //使其自加载这个文件
			})
		}
		writer.Write([]byte("back home,my son"))
	})
}
