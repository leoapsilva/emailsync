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

	listContacts, errorResponse := mockapi.GetListContacts()
	if errorResponse != nil {
		return c.JSON(errorResponse.Status, errorResponse)
	}

	syncResponse = usecases.AddContacts(listContacts)

	response, err := json.Marshal(syncResponse)
	if err != nil {
		errorResponse := model.SetErrorResponse(err.Error(), http.StatusInternalServerError, "Error Marshal Response")
		return c.JSON(http.StatusInternalServerError, errorResponse)
	}

	return c.JSON(http.StatusOK, json.RawMessage(response))
}
