package initialize

func Initialize() {
	initViper()
	initRedis()
	initGorm()
	initRouter()
}
