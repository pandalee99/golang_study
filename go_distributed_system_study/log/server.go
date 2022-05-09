package log

//因为标准库也有一个log，所以可以起一个别名
import (
	"io/ioutil"
	stlog "log"
	"net/http"
	"os"
)

var log *stlog.Logger

//目的在于把日志写入文件系统
type fileLog string

func (fl fileLog) Write(data []byte) (int, error) {
	//首先要打开文件，才能写入
	//fl文件路径，os...表示没有则创造，只写，只附加，
	f, err := os.OpenFile(string(fl), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return 0, err
	}
	//defer表示程序最后一定会执行的，这句的意思是必须把文件关闭
	defer f.Close()
	return f.Write(data)
}

//最后把log指向某个文件地址
func Run(destination string) {
	log = stlog.New(fileLog(destination), "[go]： ", stlog.LstdFlags)
}

//注册一个Handler
func RegisterHandler() {
	http.HandleFunc("/log", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost: //如果请求是post
			msg, err := ioutil.ReadAll(request.Body) //先读取内容
			if err != nil || len(msg) == 0 {         //如果有错误
				writer.WriteHeader(http.StatusBadRequest)
			}
			//没有错误则写入日志
			write(string(msg))
		default: //对于其他情况，方法就不进行，直接返回
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
}

func write(message string) {
	log.Printf("%v\n", message)
}
