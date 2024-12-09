package tests

import (
	"emailsync/model"
	"emailsync/usecases"
	"encoding/json"
	"testing"
)

func Test_CheckSyncAddingContactsTwice(t *testing.T) {

	var listContacts model.ListContacts
	var syncResponse model.SyncResponse
	var err error

	StartServer()

	SetupClient()
	t.Run("Checking /contacts/sync adding contacts twice ", func(t *testing.T) {

		responseSyncEndpoint, err = GetSyncEndpoint()
		if err != nil {
			t.Error(err)
		}

		err = json.Unmarshal(responseSyncEndpoint, &syncResponse)
		if err != nil {
			t.Error(err)
		}

		if len(syncResponse.Contacts) != 24 {
			t.Errorf("SyncedContacts: expected 24, got %d", len(syncResponse.Contacts))
		}

		listContacts = syncResponse.Contacts
		checkAddContactsResponse := usecases.AddContacts(&listContacts)

		if checkAddContactsResponse.SyncedContacts != 0 {
			t.Errorf("SyncedContacts: expected 0, got %d", checkAddContactsResponse.SyncedContacts)
		}

		t.Cleanup(ArchiveContacts)
	})
}
