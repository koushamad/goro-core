package request

import (
	"github.com/gin-gonic/gin"
	"github.com/koushamad/goro/app/exception/throw"
	"sync"
)

type (
	request interface {
		Json(*interface{})
	}
	Request struct {
		Context *gin.Context
	}
)

var (
	engine *gin.Engine
	once   sync.Once
	self   *Request
	Throw  *throw.Throw
)

func Boot(egn *gin.Engine) {
	once.Do(func() {
		engine = egn
		self = &Request{}
	})
}

func Init(ctx *gin.Context)  {
	self = &Request{}
	Throw = throw.Load()
	self.Context = ctx
}

func Load() *Request {
	return self
}

func (req Request) Get(key string) string {
	return req.Context.Query(key)
}

func (req Request) Post(key string) string {
	value, exists := req.Context.GetPostForm(key)
	if exists {
		return value
	}
	return ""
}

func (req Request) Pars(obj interface{}) {
	err := req.Context.ShouldBindJSON(obj)
	Throw.Error("on valid request", 402, err)
}
