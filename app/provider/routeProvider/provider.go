package routeProvider

import (
	"github.com/koushamad/goro-app/app/provider/routeProvider"
	"github.com/koushamad/goro-core/app/http/router"
)

func Routes(Route *router.Router) {
	routeProvider.Routes(Route)
}
