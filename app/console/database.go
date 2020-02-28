package console

import (
	database2 "github.com/koushamad/goro-app/database"
)

func database(Command string, args ...[]string) {
	switch Command {
	case "migrate":
		database2.Migration()
		break
	}
}
