package controller

import (
	"emailsync/controller/contacts"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func ConfigServerEndpoints(g *echo.Group) {
	log.Info("ConfigServerEndpoints")
	contacts.ConfigServerEndpoints(g)
	// For OpenAPI Swagger
	g.Static("/doc", "./doc")
}
