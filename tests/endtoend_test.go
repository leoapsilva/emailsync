package tests

import (
	"testing"
)

func Test_SyncEndpoint(t *testing.T) {
	var err error

	StartServer()

	responseSyncEndpoint, err = GetSyncEndpoint()

	if err != nil {
		t.Errorf("Sync %v", err)
	}

	t.Cleanup(ArchiveContacts)
}
