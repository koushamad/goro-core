package appProvider

import (
	"github.com/gin-gonic/gin"
	"github.com/koushamad/goro-app/app/provider/appProvider"
	"github.com/koushamad/goro-core/app/console"
	"github.com/koushamad/goro-core/app/exception/throw"
	"github.com/koushamad/goro-core/app/http"
	"github.com/koushamad/goro-core/app/http/request"
	"github.com/koushamad/goro-core/app/http/response"
	"github.com/koushamad/goro-core/app/http/router"
	routeProvider2 "github.com/koushamad/goro-core/app/provider/routeProvider"
)

func Boot(egn *gin.Engine) {
	response.Boot(egn)
	throw.Boot(egn)
	request.Boot(egn)
	console.Boot(egn)
	router.Boot(egn)
	http.Boot(egn)

	appProvider.Boot(egn)
}

func Run() {
	routeProvider2.Routes(router.Load())
}

func Kill() {
	appProvider.Kill()
}
