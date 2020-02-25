package console

import (
	"github.com/koushamad/goro-core/app/conf"
	"github.com/koushamad/goro-core/app/http"
)

func Serve(Command string, args []string) {
	if Command == "" {
		Command = conf.Get("app.port")
	}
	http.Load().Serve(":" + Command)
}
