package router

import "github.com/labstack/echo/v5"

var GroupRouterHubApp = &routerHub{}

type registerFunc func(public, private *echo.Group)

type routerHub struct {
	routerHubList []registerFunc
}

func (r *routerHub) InitRouterHub(public, private *echo.Group) {
	// 这里填入需要注册的路由函数
	initUserRouter()
	for _, registerFunc := range r.routerHubList {
		registerFunc(public, private)
	}
}

func (r *routerHub) RegisterRouterHub(registerFunc registerFunc) {
	r.routerHubList = append(r.routerHubList, registerFunc)
}
