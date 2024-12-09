package usecases

import (
	"emailsync/model"
	"emailsync/service/mailchimp"

	log "github.com/sirupsen/logrus"
)

func AddContacts(listContacts *model.ListContacts) *model.SyncResponse {

	var syncedContacts model.ListContacts
	var syncResponse model.SyncResponse

	for _, contact := range *listContacts {
		addedContact, errorResponse := mailchimp.AddContact(&contact)

		if errorResponse != nil {
			log.Errorf("Error on add contact: %s", errorResponse.Detail)
		} else {
			syncedContacts = append(syncedContacts, *addedContact)
		}
	}

	log.Infof("Synced [%d] from total of [%d]", len(syncedContacts), len(*listContacts))

	syncResponse.SyncedContacts = len(syncedContacts)
	syncResponse.Contacts = syncedContacts

	return &syncResponse
}
