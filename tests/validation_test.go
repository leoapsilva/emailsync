package tests

import (
	"emailsync/model"
	"encoding/json"
	"net/http"
	"testing"
)

func Test_SyncWithPOST(t *testing.T) {

	StartServer()

	SetupClient()

	t.Run("POST /contacts/sync", func(t *testing.T) {
		response, err := localAPI.Post("/contacts/sync", []byte(`{"a": "b"}`))
		if response != nil {
			if err != nil {
				errorResponse := model.GetErrorResponse(response)
				if errorResponse.Status != http.StatusMethodNotAllowed {
					t.Errorf("POST - Expected response error with HTTP Status %d got %d - %s. %s", http.StatusMethodNotAllowed, errorResponse.Status, errorResponse.Title, errorResponse.Detail)
				}
			} else {
				t.Errorf("POST - Expected %d.", http.StatusMethodNotAllowed)
			}
		} else {
			t.Errorf("POST - Expected response error with HTTP Status %d got error: %s", http.StatusMethodNotAllowed, err)

		}
	})

}

func Test_SyncWithDELETE(t *testing.T) {

	StartServer()

	SetupClient()

	t.Run("DELETE /contacts/sync", func(t *testing.T) {
		response, err := localAPI.Delete("/contacts/sync", map[string]string{}, map[string]string{})
		if response != nil {
			if err != nil {
				errorResponse := model.GetErrorResponse(response)
				if errorResponse.Status != http.StatusMethodNotAllowed {
					t.Errorf("DELETE - Expected response error with HTTP Status %d got %d - %s. %s", http.StatusMethodNotAllowed, errorResponse.Status, errorResponse.Title, errorResponse.Detail)
				}
			} else {
				t.Errorf("DELETE - Expected %d.", http.StatusMethodNotAllowed)
			}
		} else {
			t.Errorf("DELETE - Expected response error with HTTP Status %d got error: %s", http.StatusMethodNotAllowed, err)

		}
	})

}

func Test_SyncWithPUT(t *testing.T) {

	StartServer()

	SetupClient()

	t.Run("PUT /contacts/sync", func(t *testing.T) {
		response, err := localAPI.Put("/contacts/sync", []byte(`{"a": "b"}`), map[string]string{}, map[string]string{})
		if response != nil {
			if err != nil {
				errorResponse := model.GetErrorResponse(response)
				if errorResponse.Status != http.StatusMethodNotAllowed {
					t.Errorf("PUT - Expected response error with HTTP Status %d got %d - %s. %s", http.StatusMethodNotAllowed, errorResponse.Status, errorResponse.Title, errorResponse.Detail)
				}
			} else {
				t.Errorf("PUT - Expected %d.", http.StatusMethodNotAllowed)
			}
		} else {
			t.Errorf("PUT - Expected response error with HTTP Status %d got error: %s", http.StatusMethodNotAllowed, err)

		}
	})
}

func Test_SyncWithPUTEmptyPayload(t *testing.T) {

	StartServer()

	SetupClient()

	t.Run("PUT /contacts/sync empty payload", func(t *testing.T) {
		response, err := localAPI.Put("/contacts/sync", json.RawMessage{}, map[string]string{}, map[string]string{})
		if response != nil {
			if err != nil {
				errorResponse := model.GetErrorResponse(response)
				if errorResponse.Status != http.StatusBadRequest {
					t.Errorf("PUT - Expected response error with HTTP Status %d got %d - %s. %s", http.StatusBadRequest, errorResponse.Status, errorResponse.Title, errorResponse.Detail)
				}
			} else {
				t.Errorf("PUT - Expected %d.", http.StatusBadRequest)
			}
		} else {
			t.Errorf("PUT - Expected response error with HTTP Status %d got error: %s", http.StatusBadRequest, err)

		}
	})
}
