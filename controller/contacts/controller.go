package contacts

import (
	"emailsync/model"
	"emailsync/service/mailchimp"
	"emailsync/service/mockapi"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func Sync(c echo.Context) error {
	log.Info("Sync")

	var syncResponse model.SyncResponse
	var setDifference *model.MapContacts

	mockAPIMapContacts, err := mockapi.GetMapContacts()
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusServiceUnavailable, nil)
	}

	mailChimpMapContacts, err := mailchimp.GetMapContacts()
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusServiceUnavailable, nil)
	}

	setDifference = mockAPIMapContacts.SetDifference(mailChimpMapContacts)

	syncResponse.SyncedContacts = setDifference.Length()
	for _, contact := range *setDifference {
		syncResponse.Contacts = append(syncResponse.Contacts, contact)
	}

	response, err := json.Marshal(syncResponse)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusServiceUnavailable, nil)
	}

	return c.JSON(http.StatusOK, json.RawMessage(response))
}
