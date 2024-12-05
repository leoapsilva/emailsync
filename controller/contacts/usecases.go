package contacts

import (
	"emailsync/model"
	"emailsync/service/mailchimp"

	"github.com/labstack/gommon/log"
)

func addContact(mockAPIMapContacts *model.MapContacts) *model.SyncResponse {

	var syncedContacts model.MapContacts
	var syncResponse model.SyncResponse

	syncedContacts = *mockAPIMapContacts

	added := 0
	for _, contact := range *mockAPIMapContacts {
		errorResponse := mailchimp.AddContact(&contact)

		if errorResponse != nil {
			log.Errorf("Error on add contact: %s", errorResponse.Detail)
			delete(syncedContacts, contact.Email)
		} else {
			added = added + 1
		}
	}

	log.Infof("Synced [%d] from total of [%d]", added, mockAPIMapContacts.Length())

	syncResponse.SyncedContacts = added
	for _, contact := range syncedContacts {
		syncResponse.Contacts = append(syncResponse.Contacts, contact)
	}

	return &syncResponse
}
