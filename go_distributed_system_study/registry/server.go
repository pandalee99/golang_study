package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

const ServerPort = ":3000"
const ServicesURL = "http://localhost" + ServerPort + "/services"

type registry struct {
	//这个slice，是动态变化的，而且多个线程可能会并发的进行访问，
	//为了保证线程安全，需要加锁。
	registrations []Registration
	mutex         *sync.RWMutex
}

//在增加服务的时候是需要加锁的
//add方法表示增加服务
func (registry *registry) add(reg Registration) error {
	registry.mutex.Lock()
	registry.registrations = append(registry.registrations, reg)
	registry.mutex.Unlock()

	//正好去获取这个服务所依赖的其他服务
	err := registry.sendRequireServices(reg)

	//服务通知，当服务上线，而这个服务又被依赖时，告知依赖服务自己上线了
	registry.notify(patch{
		Added: []patchEntry{
			patchEntry{
				Name: reg.ServiceName,
				URL:  reg.ServiceURL,
			},
		},
	})

	return err
}

//获取其他服务
func (registry *registry) sendRequireServices(reg Registration) error {
	registry.mutex.RLock()
	defer registry.mutex.RUnlock()

	//寻找服务
	var p patch
	for _, serviceReg := range registry.registrations {
		for _, reqService := range reg.RequiredServices {
			if serviceReg.ServiceName == reqService {
				p.Added = append(p.Added, patchEntry{
					Name: serviceReg.ServiceName,
					URL:  serviceReg.ServiceURL,
				})
			}
		}
	}
	//找到之后,注册
	err := registry.sendPatch(p, reg.ServiceUpdateURL)
	if err != nil {
		return err
	}
	return nil
}

//通知其他服务
func (r registry) notify(fullPatch patch) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	//看看服务的依赖在patch里面存不存在
	//对已经注册的服务循环遍历
	for _, reg := range r.registrations {
		//并发的发出通知
		go func(reg Registration) {
			//对服务所需要的服务进行循环
			for _, reqService := range reg.RequiredServices {
				p := patch{Added: []patchEntry{}, Removed: []patchEntry{}}
				//标志位，为TRUE表示有需要更新的地方
				sendUpdate := false
				for _, added := range fullPatch.Added {
					//如果添加的服务正好是某个服务的依赖项
					if added.Name == reqService {
						p.Added = append(p.Added, added)
						sendUpdate = true
					}
				}
				//看看有哪些服务停止了
				for _, removed := range fullPatch.Removed {
					///如果停掉的服务正好是所被依赖的服务
					if removed.Name == reqService {
						p.Removed = append(p.Removed, removed)
						sendUpdate = true
					}
				}
				//最后判断标志位，把更新发送到对应的服务
				if sendUpdate {
					err := r.sendPatch(p, reg.ServiceUpdateURL)
					if err != nil {
						log.Println(err)
						return
					}
				}

			}
		}(reg)
	}
}

//将需要的服务发送过去注册的过程
func (r registry) sendPatch(p patch, url string) error {
	d, err := json.Marshal(p)
	if err != nil {
		return err
	}
	_, err = http.Post(url, "application/json", bytes.NewBuffer(d))
	if err != nil {
		return err
	}
	return nil
}

//那么肯定要有个方法取消服务
func (registry *registry) remove(url string) error {
	//看看服务中台有没有这服务
	for i := range reg.registrations {
		if reg.registrations[i].ServiceURL == url {

			//下线也告知
			registry.notify(patch{
				Removed: []patchEntry{
					{
						Name: registry.registrations[i].ServiceName,
						URL:  registry.registrations[i].ServiceURL,
					},
				},
			})
			//下线告知

			registry.mutex.Lock()

			//把后面的接上前面的，自然的去除了，方法其实不唯一
			//这个写的不对，但是运行对
			reg.registrations = append(reg.registrations[:i], reg.registrations[:i+1]...)
			//这个写的对，但运行会出现连续两次remove，导致看起来不对
			//reg.registrations = append(reg.registrations[:i], reg.registrations[i+1:]...)
			registry.mutex.Unlock()
			return nil
		}
	}
	return fmt.Errorf("服务地址未发现： %s ", url)
}

var reg = registry{
	registrations: make([]Registration, 0),
	mutex:         new(sync.RWMutex),
}

//创建一个web服务
type RegistryService struct{}

func (s RegistryService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("等待接受请求：")
	switch r.Method {
	case http.MethodPost:
		//解码
		dec := json.NewDecoder(r.Body)
		var r Registration
		err := dec.Decode(&r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Printf("增加服务：%v ，该服务的地址是：%s \n",
			r.ServiceName, r.ServiceURL)
		err = reg.add(r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	case http.MethodDelete:
		payload, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		url := string(payload)
		log.Printf("移除服务： %s", url)
		err = reg.remove(url)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

var once sync.Once

func SetupRegistryService() {
	once.Do(func() {
		go reg.heartbeat(3 * time.Second)
	})
}

//每隔一段时间，进行监测检查
func (r *registry) heartbeat(freq time.Duration) {
	for {
		var wg sync.WaitGroup
		for _, reg := range r.registrations {
			wg.Add(1)
			go func(reg Registration) {
				defer wg.Done()
				success := true
				for attemps := 0; attemps < 3; attemps++ {
					res, err := http.Get(reg.HeartbeatURL)
					if err != nil {
						log.Println(err)
					} else if res.StatusCode == http.StatusOK {
						log.Printf("Heartbeat check passed for %v", reg.ServiceName)
						if !success {
							r.add(reg)
						}
						break
					}
					log.Printf("Heartbeat check failed for %v", reg.ServiceName)
					if success {
						success = false
						r.remove(reg.ServiceURL)
					}
					time.Sleep(1 * time.Second)
				}
			}(reg)
			wg.Wait()
			time.Sleep(freq)
		}
	}
}
