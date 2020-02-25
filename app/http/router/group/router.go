package group

import (
	"github.com/gin-gonic/gin"
	"github.com/koushamad/goro/app/http/middleware"
	"github.com/koushamad/goro/app/http/request"
	"github.com/koushamad/goro/app/http/response"
	contextProvider2 "github.com/koushamad/goro/app/provider/contextProvider"
)

type (
	Router struct {
		RouteGroup *gin.RouterGroup
	}
)

func Load(routeGroup *gin.RouterGroup) *Router {
	return &Router{routeGroup}
}

func (router Router) Middleware(mid func(mid *middleware.Middleware)) {
	router.RouteGroup.Use(func(ctx *gin.Context) {
		contextProvider2.Init(ctx)
		mid(middleware.Load())
	})
}

func (router *Router) GroupMiddleware(path string, mid func(mid *middleware.Middleware), handler func(router *Router)) {
	Router := router.RouteGroup.Group(path)
	Router.Use(func(ctx *gin.Context) {
		contextProvider2.Init(ctx)
		mid(middleware.Load())
	})
	{
		handler(Load(Router))
	}
}

func (router *Router) Group(path string, handler func(router *Router)) {
	Router := router.RouteGroup.Group(path)
	Router.Use(func(ctx *gin.Context) {
		contextProvider2.Init(ctx)
	})
	{
		handler(Load(Router))
	}
}

func (router Router) Get(path string, controller func(req *request.Request, res *response.Response)) {
	router.RouteGroup.GET(path, func(ctx *gin.Context) {
		contextProvider2.Init(ctx)
		controller(request.Load(), response.Load())
	})
}

func (router Router) Post(path string, controller func(req *request.Request, res *response.Response)) {
	router.RouteGroup.POST(path, func(ctx *gin.Context) {
		contextProvider2.Init(ctx)
		controller(request.Load(), response.Load())
	})
}

func (router Router) Put(path string, controller func(req *request.Request, res *response.Response)) {
	router.RouteGroup.PUT(path, func(ctx *gin.Context) {
		contextProvider2.Init(ctx)
		controller(request.Load(), response.Load())
	})
}

func (router Router) Patch(path string, controller func(req *request.Request, res *response.Response)) {
	router.RouteGroup.PATCH(path, func(ctx *gin.Context) {
		contextProvider2.Init(ctx)
		controller(request.Load(), response.Load())
	})
}

func (router Router) Delete(path string, controller func(req *request.Request, res *response.Response)) {
	router.RouteGroup.DELETE(path, func(ctx *gin.Context) {
		contextProvider2.Init(ctx)
		controller(request.Load(), response.Load())
	})
}
