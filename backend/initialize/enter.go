package initialize

func Initialize() {
	initViper()
	initRedis()
	InitEmail()
	initGorm()
	initRouter()
}
