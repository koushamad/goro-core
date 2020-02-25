package goro_core

import (
	"github.com/gin-gonic/gin"
	"github.com/koushamad/goro-core/app/http"
	"github.com/koushamad/goro-core/app/log"
	"github.com/koushamad/goro-core/app/provider/appProvider"
	"github.com/koushamad/goro-core/app/provider/configProvider"
)

func Boot() {
	configProvider.Load()
	appProvider.Boot(engine())
	appProvider.Run()
	http.Load().Serve(":8000")
}

func Kill()  {
	appProvider.Kill()
}

func engine() *gin.Engine {
	egn := gin.Default()
	egn.Use(log.Logger())
	egn.Use(gin.Recovery())
	return egn
}
