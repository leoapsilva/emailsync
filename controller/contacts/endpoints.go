package contacts

import (
	"github.com/labstack/echo/v4"
)

const (
	RouteContactsSync string = "/contacts/sync"
)

func ConfigServerEndpoints(g *echo.Group) {
	g.GET(RouteContactsSync, Sync)
}
