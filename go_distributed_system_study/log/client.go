package log

import (
	"bytes"
	"fmt"
	"go_distributed_system_study/registry"
	"net/http"

	stlog "log"
)

//写日志，把日志写到server
func SetClientLogger(serviceURL string, clientService registry.ServiceName) {
	stlog.SetPrefix(fmt.Sprintf("[%v] - ", clientService))
	stlog.SetFlags(0)
	stlog.SetOutput(&clientLogger{url: serviceURL})
}

type clientLogger struct {
	url string
}

func (cl clientLogger) Write(data []byte) (int, error) {
	b := bytes.NewBuffer([]byte(data))
	//写到服务端
	res, err := http.Post(cl.url+"/log", "text/plain", b)
	if err != nil {
		return 0, err
	}
	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("Failed to send log message. Service responded with %d - %s", res.StatusCode, res.Status)
	}
	//如果都没有问题，返回数据
	return len(data), nil
}
