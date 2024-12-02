package contacts

import (
	"emailsync/model"
	"emailsync/service/mockapi"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func Sync(c echo.Context) error {
	log.Info("Sync")

	var syncResponse model.SyncResponse

	mockAPIMapContacts, err := mockapi.GetListContacts()
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusServiceUnavailable, nil)
	}

	syncResponse.SyncedContacts = mockAPIMapContacts.Length()
	for _, contact := range *mockAPIMapContacts {
		syncResponse.Contacts = append(syncResponse.Contacts, contact)
	}

	response, err := json.Marshal(syncResponse)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusServiceUnavailable, nil)
	}

	return c.JSON(http.StatusOK, response)
}
