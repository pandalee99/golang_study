package main


/*
func main() {
	controller.RegisterRoutes()
	http.ListenAndServeTLS("localhost:8888", "cert.pem", "key.pem", nil)

}
*/
/*
type Company struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

func main() {
	http.HandleFunc("/companies", func(writer http.ResponseWriter, request *http.Request) {
		c := Company{
			ID:      111,
			Name:    "Microsoft",
			Country: "USA",
		}
		//time.Sleep(4*time.Second)
		//测试用
		enc := json.NewEncoder(writer)
		enc.Encode(c)
	})
	http.ListenAndServe("localhost:8888", &middleware.TimeoutMiddleware{Next: new(middleware.AuthMiddleware)})

}
*/
/*
type helloHandler struct {
}

//http.ResponseWriter用于表达响应的
// *http.Request 收到的请求，可以是get或post等等
func (m *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, my lover"))
}

type aboutHandler struct {
}

func (m *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About, my message"))
}

func welcome(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("welcome,my friends"))
}

func main() {
	//自制一个Handler
	mh := helloHandler{}
	a := aboutHandler{}

	//等同于 http.ListenAndServe
	server := http.Server{
		Addr:    "localhost:8888",
		Handler: nil, //这样会默认使用DefaultServeMux
		//http.DefaultServeMux也是一个Handler，使用nil是默认使用
	}
	http.Handle("/hello", &mh) //相信到这里，就很简单了，直接望文生义即可
	http.Handle("/about", &a)

	//使用HandleFunc
	http.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("back home,my son"))
	})
	http.HandleFunc("/welcome", http.HandlerFunc(welcome))
	server.ListenAndServe()

}
*/
/*
type Company struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

func main() {

	http.HandleFunc("/companies", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost: //首先看看请求的类型
			dec := json.NewDecoder(request.Body) //现在JSON就在这个body里边了，并且需要进行解码
			company := Company{}
			err := dec.Decode(&company) //将密码解码到company这个变量里边
			if err != nil {             //如果它不为空
				log.Println(err.Error())                           //打印错误
				writer.WriteHeader(http.StatusInternalServerError) //500error
				return
			}
			//否则将客户端传过来的数据转化为json重新返回给客户端
			enc := json.NewEncoder(writer)
			err = enc.Encode(company)
			if err != nil {
				log.Println(err.Error())                           //打印错误
				writer.WriteHeader(http.StatusInternalServerError) //500error
				return
			}
		default:
			writer.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	server := http.Server{
		Addr:    "localhost:8888",
		Handler: nil,
	}
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
*/
