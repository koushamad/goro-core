package console

import (
	"github.com/fatih/color"
	DB "github.com/koushamad/goro-app/database"
)

func database(Command string, args ...[]string) {
	switch Command {
	case "migrate":
		DB.Migration()
		color.Green("Database migration done")
		break
	}
}
