package contextProvider

import (
	"github.com/gin-gonic/gin"
	"github.com/koushamad/goro-app/app/provider/contextProvider"
	"github.com/koushamad/goro/app/exception/throw"
	"github.com/koushamad/goro/app/http/middleware"
	"github.com/koushamad/goro/app/http/request"
	"github.com/koushamad/goro/app/http/response"
)

func Init(ctx *gin.Context) {
	response.Init(ctx)
	throw.Init(ctx)
	middleware.Init(ctx)
	request.Init(ctx)

	contextProvider.Init(ctx)
}