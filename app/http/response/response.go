package response

import (
	"github.com/gin-gonic/gin"
	"sync"
)

type (
	Response struct {
		Context *gin.Context
		Accept  string
		res     interface{}
	}
)

var (
	once   sync.Once
	engine *gin.Engine
	self   *Response
)

func Boot(egn *gin.Engine) {
	once.Do(func() {
		engine = egn
	})
}

func Init(ctx *gin.Context) {
	self = &Response{Context:ctx}
}

func Load() *Response {
	return self
}

func (res Response) Json(code int, response interface{}) {
	if !res.Context.IsAborted() {
		res.Context.Abort()
		res.Context.JSON(code, response)
	}
}

func (res Response) Xml(code int, response interface{}) {
	if !res.Context.IsAborted() {
		res.Context.Abort()
		res.Context.XML(code, response)
	}
}

func (res Response) Send(code int, response interface{}) {
	switch res.Accept {
	case "application/xml":
	case "application/atom+xml":
	case "application/rdf+xml":
	case "application/rss+xml":
	case "application/soap+xml":
	case "application/xhtml+xml":
	case "application/xml-dtd":
	case "application/xop+xml":
	case "image/svg+xml":
	case "message/imdn+xml":
	case "text/xml":
		res.Xml(code, response)
		break
	case "application/json":
	default:
		res.Json(code, response)
	}
}
