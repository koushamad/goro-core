package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/koushamad/goro/app/exception/throw"
	"sync"
)

type (
	Middleware struct {
		Context *gin.Context
		Throw *throw.Throw
	}
)

var (
	engine *gin.Engine
	once   sync.Once
	self   *Middleware
)

func Boot(egn *gin.Engine) {
	once.Do(func() {
		engine = egn
		self = &Middleware{}
	})
}

func Init(ctx *gin.Context) *Middleware {
	self = &Middleware{
		Context: ctx,
		Throw: throw.Load(),
	}
	return self
}

func Load() *Middleware {
	return self
}

func (mid *Middleware) Set(key string, value interface{}) {
	mid.Context.Set(key, value)
}

func (mid *Middleware) Next() {
	mid.Context.Next()
}

func (mid *Middleware) Abort() {
	mid.Context.Abort()
}