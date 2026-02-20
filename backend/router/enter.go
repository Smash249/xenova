package router

import "github.com/labstack/echo/v5"

var GroupRouterHubApp = &routerHub{}

type registerFunc func(public, private, admin *echo.Group)

type routerHub struct {
	routerHubList []registerFunc
}

func (r *routerHub) InitRouterHub(public, private, admin *echo.Group) {
	initUserRouter()
	initProductRouter()
	initJournalism()
	initSoftwareRouter()
	initHonorRouter()
	for _, registerFunc := range r.routerHubList {
		registerFunc(public, private, admin)
	}
}

func (r *routerHub) RegisterRouterHub(registerFunc registerFunc) {
	r.routerHubList = append(r.routerHubList, registerFunc)
}
