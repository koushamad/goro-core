package router

import (
	"github.com/gin-gonic/gin"
	"github.com/koushamad/goro/app/http/middleware"
	"github.com/koushamad/goro/app/http/request"
	"github.com/koushamad/goro/app/http/response"
	"github.com/koushamad/goro/app/http/router/group"
	contextProvider2 "github.com/koushamad/goro/app/provider/contextProvider"
	"sync"
)

type (
	Router struct{}
)

var (
	once   sync.Once
	self   *Router
	engine *gin.Engine
)

func Boot(egn *gin.Engine) {
	once.Do(func() {
		engine = egn
		Init()
	})
}

func Init() {
	self = &Router{}
}

func Load() *Router {
	return self
}

func (router Router) Middleware(mid func(mid *middleware.Middleware)) {
	engine.Use(func(ctx *gin.Context) {
		contextProvider2.Init(ctx)
		mid(middleware.Load())
	})
}

func (router *Router) GroupMiddleware(path string, mid func(mid *middleware.Middleware),
	handler func(router *group.Router)) {
	Router := engine.Group(path)
	Router.Use(func(ctx *gin.Context) {
		contextProvider2.Init(ctx)
		mid(middleware.Load())
	})
	{
		handler(group.Load(Router))
	}
}

func (router *Router) Group(path string, handler func(router *group.Router)) {
	Router := engine.Group(path)
	Router.Use(func(ctx *gin.Context) {
		contextProvider2.Init(ctx)
	})
	{
		handler(group.Load(Router))
	}
}

func (router Router) Get(path string, controller func(req *request.Request, res *response.Response)) {
	engine.GET(path, func(ctx *gin.Context) {
		contextProvider2.Init(ctx)
		controller(request.Load(), response.Load())
	})
}

func (router Router) Post(path string, controller func(req *request.Request, res *response.Response)) {
	engine.POST(path, func(ctx *gin.Context) {
		contextProvider2.Init(ctx)
		controller(request.Load(), response.Load())
	})
}

func (router Router) Put(path string, controller func(req *request.Request, res *response.Response)) {
	engine.PUT(path, func(ctx *gin.Context) {
		contextProvider2.Init(ctx)
		controller(request.Load(), response.Load())
	})
}

func (router Router) Patch(path string, controller func(req *request.Request, res *response.Response)) {
	engine.PATCH(path, func(ctx *gin.Context) {
		contextProvider2.Init(ctx)
		controller(request.Load(), response.Load())
	})
}

func (router Router) Delete(path string, controller func(req *request.Request, res *response.Response)) {
	engine.DELETE(path, func(ctx *gin.Context) {
		contextProvider2.Init(ctx)
		controller(request.Load(), response.Load())
	})
}

func (router Router)Any(controller func(req *request.Request, res *response.Response))  {
	engine.NoRoute(func(ctx *gin.Context) {
		contextProvider2.Init(ctx)
		controller(request.Load(), response.Load())
	})
}
