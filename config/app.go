package config

import "github.com/koushamad/goro-core/app/helper"

var App = map[string]string{
	"port": helper.Env("APP_PORT", "8000"),
	"debug": helper.Env("APP_DEBUG", "develop"),
}