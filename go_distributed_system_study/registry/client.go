package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
)

//这个函数的目的是给web service发送一个post请求
func RegisterService(r Registration) error {

	heartbeatURL, err := url.Parse(r.HeartbeatURL)
	if err != nil {
		return err
	}
	http.HandleFunc(heartbeatURL.Path, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	//服务注册中心要向URL更新一些信息
	serviceUpdateURL, err := url.Parse(r.ServiceUpdateURL)
	if err != nil {
		return err
	}
	http.Handle(serviceUpdateURL.Path, &serviceUpdateHandler{})
	//

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err = enc.Encode(r)
	if err != nil {
		return err
	}

	res, err := http.Post(ServicesURL, "application/json", buf)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("服务注册失败 "+"状态码为： %v", res.StatusCode)
	}
	return nil
}

//更新服务的处理
type serviceUpdateHandler struct{}

func (suh serviceUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	//先进行解码
	dec := json.NewDecoder(r.Body)
	var p patch
	err := dec.Decode(&p)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("收到更新： %v\n", p)

	prov.Update(p)
}

//结束服务
func ShutdownService(url string) error {
	req, err := http.NewRequest(
		http.MethodDelete, ServicesURL,
		bytes.NewBuffer([]byte(url))) //把string转化为slice
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "text/plain")
	//紧接着发送请求
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	//fmt.Println(strconv.Itoa(res.StatusCode))
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("服务取消失败，状态码为：%v",
			res.StatusCode)
	}
	return nil
}

//log服务会向多个服务提供服务
type providers struct {
	//服务与服务的URL
	services map[ServiceName][]string
	//互斥锁
	mutex *sync.RWMutex
}

//接受到patch的时候，需要进行更新，
var prov = providers{
	services: make(map[ServiceName][]string),
	mutex:    new(sync.RWMutex),
}

//实现
func (p *providers) Update(pat patch) {
	//对传进来的patch更新provider
	p.mutex.Lock()
	defer p.mutex.Unlock()

	//新增的情况
	for _, patchEntry := range pat.Added {
		//如果这个服务名目前还不存在，就创建新的slice
		if _, ok := p.services[patchEntry.Name]; !ok {
			p.services[patchEntry.Name] = make([]string, 0)
		}
		//如果存在的话，就在值后边附加URL
		p.services[patchEntry.Name] = append(p.services[patchEntry.Name],
			patchEntry.URL)
	}
	//减少的情况
	//遍历，对比，移除
	for _, patchEntry := range pat.Removed {
		if providerURLs, ok := p.services[patchEntry.Name]; ok {
			for i := range providerURLs {
				if providerURLs[i] == patchEntry.URL {
					p.services[patchEntry.Name] = append(providerURLs[:i],
						providerURLs[i+1:]...)
				}
			}
		}
	}
}

//然后还需要，根据服务的名称来找到它所依赖服务的url
func (p providers) get(name ServiceName) (string, error) {
	providers, ok := p.services[name]
	if !ok {
		return "", fmt.Errorf("没有可提供服务的提供商： %v", name)
	}
	//随机数
	idx := int(rand.Float32() * float32(len(providers)))
	return providers[idx], nil
}

//由于这个get方法是私有的，对外再套一个函数：
func GetProvider(name ServiceName) (string, error) {
	return prov.get(name)
}
