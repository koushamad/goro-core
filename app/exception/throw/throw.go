package throw

import (
	"github.com/gin-gonic/gin"
	exceptionHandler "github.com/koushamad/goro-app/app/exception"
	"github.com/koushamad/goro-app/app/http/contrroller/resource"
	"github.com/koushamad/goro/app/exception"
	"github.com/koushamad/goro/app/http/response"
	"github.com/labstack/gommon/log"
	"sync"
)

type (
	Throw struct {
		Context *gin.Context
		Response *response.Response
	}
)

var (
	engine *gin.Engine
	once   sync.Once
	self Throw
)

func Boot(egn *gin.Engine) {
	once.Do(func() {
		engine = egn
	})
}

func Init(ctx *gin.Context) {
	self = Throw{Context:ctx, Response: response.Load()}
}

func Load() *Throw {
	return &self
}

func (throw Throw)handler(e exception.Exception)  {
	statusCode, data := exceptionHandler.Handler(e)
	if statusCode >= 300 {
		if throw.Context != nil {
			if data != nil {
				throw.Response.Send(statusCode, data)
			} else {
				throw.Response.Send(statusCode, resource.Error(e))
			}
		}
	}
}

func (throw Throw) Fatal(message string, code int, err error) {
	if err != nil {
		e := exception.Exception{Message: message, Code: code, Err: err}
		throw.handler(e)
		log.Fatal(e)
	}
}

func (throw Throw) Error(message string, code int, err error) {
	if err != nil {
		e := exception.Exception{Message: message, Code: code, Err: err}
		throw.handler(e)
		log.Error(e)
	}
}

func (throw Throw) Debug(message string, code int, err error) {
	if err != nil {
		e := exception.Exception{Message: message, Code: code, Err: err}
		throw.handler(e)
		log.Debug(e)
	}
}

func (throw Throw) Info(message string, code int, err error) {
	if err != nil {
		e := exception.Exception{Message: message, Code: code, Err: err}
		log.Info(e)
	}
}

func Fatal(message string, code int, err error) {
	if err != nil {
		e := exception.Exception{Message: message, Code: code, Err: err}
		exceptionHandler.Handler(e)
		log.Fatal(e.Err)
	}
}

func Error(message string, code int, err error) {
	if err != nil {
		e := exception.Exception{Message: message, Code: code, Err: err}
		exceptionHandler.Handler(e)
		log.Error(e)
	}
}

func Info(message string, code int, err error) {
	if err != nil {
		e := exception.Exception{Message: message, Code: code, Err: err}
		exceptionHandler.Handler(e)
		log.Info(e)
	}
}

func Debug(message string, code int, err error) {
	if err != nil {
		e := exception.Exception{Message: message, Code: code, Err: err}
		exceptionHandler.Handler(e)
		log.Debug(e)
	}
}