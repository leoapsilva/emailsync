package usecases

import (
	"emailsync/model"
	"emailsync/service/mailchimp"

	log "github.com/sirupsen/logrus"
)

func AddContacts(mapContacts *model.MapContacts) *model.SyncResponse {

	var syncedContacts model.MapContacts
	var syncResponse model.SyncResponse

	syncedContacts = *mapContacts

	added := 0
	for _, contact := range *mapContacts {
		errorResponse := mailchimp.AddContact(&contact)

		if errorResponse != nil {
			log.Errorf("Error on add contact: %s", errorResponse.Detail)
			delete(syncedContacts, contact.Email)
		} else {
			added = added + 1
		}
	}

	log.Infof("Synced [%d] from total of [%d]", added, mapContacts.Length())

	syncResponse.SyncedContacts = added
	for _, contact := range syncedContacts {
		syncResponse.Contacts = append(syncResponse.Contacts, contact)
	}

	return &syncResponse
}
