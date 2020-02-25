package configProvider

import (
	appConfigProvider "github.com/koushamad/goro-app/app/provider/configProvider"
	"github.com/koushamad/goro-core/app/conf"
	"github.com/koushamad/goro-core/config"
)

func Load() {
	conf.Init()

	conf.Add("App", config.App)
	conf.Add("App", config.Log)

	appConfigProvider.Load()
}
