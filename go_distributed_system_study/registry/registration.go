package registry

//注册中心
type Registration struct {
	ServiceName ServiceName
	ServiceURL  string

	//该服务所依赖的其他服务。用slice去保存
	RequiredServices []ServiceName
	//向外暴露的服务端口
	ServiceUpdateURL string
	//定时检查
	HeartbeatURL string
}

type ServiceName string

const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
	PortalService  = ServiceName("Portald")
)

//每一条目
type patchEntry struct {
	Name ServiceName
	URL  string
}

//服务变化
type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}
