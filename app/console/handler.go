package console

import "github.com/koushamad/goro-app/app/console/command"

func commands()  {
	Commends = map[string]map[string]string{
		"serve":  {"": "Listen to port"},
		"database": {"migrate": "Create database tables"},
	}
}

func handler(Row string,Command string,attr []string)  {
	switch Row {
	case "serve":
		Serve(Command, attr)
		break
	case "database":
		database(Command, attr)
		break
	}
	command.Handler(Row ,Command ,attr)
}