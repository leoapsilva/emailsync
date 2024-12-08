package tests

import (
	"emailsync/app"
	"emailsync/config"
	"emailsync/model"
	"emailsync/service"
	"emailsync/service/mailchimp"
	"encoding/json"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
)

var responseSyncEndpoint json.RawMessage

func StartServer() {
	go app.Start()
	time.Sleep(service.Timeout2s)
}

func GetSyncEndpoint() (json.RawMessage, error) {

	config.LoadEnvVariables()

	serverPort := config.GetEnvVariable("SERVER_PORT")

	localAPI := service.NewWithConnection(model.Connection{URL: "http://localhost:" + serverPort})

	return localAPI.Get("/contacts/sync")
}

func ArchiveContacts() {
	var syncResponse model.SyncResponse

	err := json.Unmarshal(responseSyncEndpoint, &syncResponse)
	if err != nil {
		log.Error(err)
		return
	}

	for _, contact := range syncResponse.Contacts {
		mailchimp.ArchiveContact(&contact)
	}
}

func Test_SyncEndpoint(t *testing.T) {
	var err error

	StartServer()

	responseSyncEndpoint, err = GetSyncEndpoint()

	if err != nil {
		t.Errorf("Sync %v", err)
	}

	t.Cleanup(ArchiveContacts)
}
