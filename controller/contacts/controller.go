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
	var syncedContacts model.MapContacts

	mockAPIMapContacts, errorResponse := mockapi.GetMapContacts()
	if errorResponse != nil {
		return c.JSON(errorResponse.Status, errorResponse)
	}

	added := 0

	setDifference = mockAPIMapContacts

	for _, contact := range *setDifference {
		errorResponse := mailchimp.AddContact(&contact)

		if errorResponse != nil {
			log.Errorf("Error on add contact: %s", errorResponse.Detail)
			delete(syncedContacts, contact.Email)
		} else {
			added = added + 1
		}
	}

	log.Infof("Synced [%d] from total of [%d]", added, setDifference.Length())

	syncResponse.SyncedContacts = added
	for _, contact := range *setDifference {
		syncResponse.Contacts = append(syncResponse.Contacts, contact)
	}

	response, err := json.Marshal(syncResponse)
	if err != nil {
		errorResponse := mailchimp.SetErrorResponse(err.Error(), http.StatusInternalServerError, "Error Marshal Response")
		return c.JSON(http.StatusInternalServerError, errorResponse)
	}

	return c.JSON(http.StatusOK, json.RawMessage(response))
}
