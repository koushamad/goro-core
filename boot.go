package goro_core

import (
	"github.com/gin-gonic/gin"
	DB "github.com/koushamad/goro-app/database"
	"github.com/koushamad/goro-core/app/log"
	"github.com/koushamad/goro-core/app/provider/appProvider"
	"github.com/koushamad/goro-core/app/provider/configProvider"
)

func Boot() {
	configProvider.Load()
	appProvider.Boot(engine())
	appProvider.Run()
	DB.Migration()
}

func Kill() {
	appProvider.Kill()
}

func engine() *gin.Engine {
	egn := gin.Default()
	egn.Use(log.Logger())
	egn.Use(gin.Recovery())
	return egn
}
