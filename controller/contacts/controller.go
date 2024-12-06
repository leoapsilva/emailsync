package contacts

import (
	"emailsync/model"
	"emailsync/service/mockapi"
	"emailsync/usecases"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func Sync(c echo.Context) error {
	log.Info("Sync")

	var syncResponse *model.SyncResponse

	mapContacts, errorResponse := mockapi.GetMapContacts()
	if errorResponse != nil {
		return c.JSON(errorResponse.Status, errorResponse)
	}

	syncResponse = usecases.AddContacts(mapContacts)

	response, err := json.Marshal(syncResponse)
	if err != nil {
		errorResponse := model.SetErrorResponse(err.Error(), http.StatusInternalServerError, "Error Marshal Response")
		return c.JSON(http.StatusInternalServerError, errorResponse)
	}

	return c.JSON(http.StatusOK, json.RawMessage(response))
}

func Add(c echo.Context) error {

	var listContacts model.ListContacts
	var syncResponse *model.SyncResponse

	err := c.Bind(&listContacts)
	if err != nil {
		errorResponse := model.SetErrorResponse(err.Error(), http.StatusInternalServerError, "Error Unmarshal bind")
		return c.JSON(http.StatusInternalServerError, errorResponse)
	}

	mapContacts := listContacts.ToMapContacts()

	syncResponse = usecases.AddContacts(mapContacts)

	response, err := json.Marshal(syncResponse)
	if err != nil {
		errorResponse := model.SetErrorResponse(err.Error(), http.StatusInternalServerError, "Error Marshal Response")
		return c.JSON(http.StatusInternalServerError, errorResponse)
	}

	return c.JSON(http.StatusOK, json.RawMessage(response))
}
