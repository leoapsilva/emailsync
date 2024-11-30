package controller

import (
	"emailsync/controller/contacts"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func ConfigEndpoints(g *echo.Group) {
	log.Info("ConfigEndpoints")
	contacts.ConfigEndpoints(g)
	// For OpenAPI Swagger
	g.Static("/doc", "./doc")
}
