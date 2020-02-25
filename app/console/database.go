package console

import "github.com/koushamad/goro-app/database/migration"

func database(Command string, args ...[]string) {
	switch Command {
	case "migrate":
		migration.Migration()
		break
	}
}

