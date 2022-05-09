package controller

//RegisterRoutes 用于注册
func RegisterRoutes() {
	//这里还可以使用静态资源
	registerHomeRoutes()
	registerWelcomeRoutes()
	RegisterRoutesController()
}
