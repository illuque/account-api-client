package client

import (
	"fmt"
	"github.com/illuque/account-api-client/model"
	"net/http"
)

func (ac AccountHttpClient) Delete(id model.DeleteId) (deleted bool, errorData *model.ErrorData) {
	ac.logger.Debugf("Calling API for Delete with id [%+v]", id)

	deleteUri := fmt.Sprintf("%s/%s?version=%d", ac.uri, id.Id, id.Version)

	// Create request
	deleteRequest, err := http.NewRequest("DELETE", deleteUri, nil)
	if err != nil {
		ac.logger.WithError(err).Errorf("Error sending DELETE to Account API")
		return
	}

	// Fetch Request
	response, err := ac.httpClient.Do(deleteRequest)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch response.StatusCode {
	case http.StatusNoContent:
		deleted = true
	case http.StatusNotFound:
		errorData = model.NewNotFound("Specified resource does not exist")
	case http.StatusConflict:
		errorData = model.NewConflict("Specified version incorrect")
	default:
		errorMsg, _ := ac.getErrorFromResponse(response)
		errorData = model.NewUnknownClientError("Unknown error code received from API on DELETE: " + errorMsg)
	}

	return
}
