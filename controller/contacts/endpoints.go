package contacts

import (
	"github.com/labstack/echo/v4"
)

const (
	RouteContactsSync string = "/contacts/sync"
	RouteContactsAdd  string = "/contacts/add"
)

func ConfigServerEndpoints(g *echo.Group) {
	g.GET(RouteContactsSync, Sync)
	g.POST(RouteContactsAdd, Add)
}
