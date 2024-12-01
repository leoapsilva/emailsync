package contacts

import (
	"emailsync/config"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func Sync(c echo.Context) error {
	log.Info("Sync")

	contactsEndpoint := config.ConfigEnvVariable("CONTACTS_ENDPOINT")
	log.Infof("Getting the contacts from %s to sync", contactsEndpoint)

	return c.JSON(http.StatusOK, nil)
}
