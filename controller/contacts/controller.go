package contacts

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func Sync(c echo.Context) error {
	log.Info("Sync")

	return c.JSON(http.StatusOK, nil)
}
