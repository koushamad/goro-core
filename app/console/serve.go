package console

import (
	"github.com/koushamad/goro/app/conf"
	"github.com/koushamad/goro/app/http"
)

func Serve(Command string, args []string)  {
	if Command ==  "" {
		Command = conf.Get("app.port")
	}
	http.Load().Serve(":" + Command)
}
