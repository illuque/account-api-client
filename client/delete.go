package client

import (
	"fmt"
	"github.com/illuque/account-api-client/model/client_error"
	"net/http"
)

// TODO:I ValueObject
func (ac AccountHttpClient) Delete(id string, version int64) (deleted bool, errorData *client_error.ErrorData) {
	ac.logger.Debugf("Calling API for Delete with id [%s] and version [%d]...", id, version)

	deleteUri := fmt.Sprintf("%s/%s?version=%d", ac.uri, id, version)

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
		errorData = client_error.NewNotFound("Specified resource does not exist")
	case http.StatusConflict:
		errorData = client_error.NewConflict("Specified version incorrect")
	default:
		errorMsg, _ := ac.getErrorFromResponse(response)
		errorData = client_error.NewUnknownClientError("Unknown error code received from API on DELETE: " + errorMsg)
	}

	return
}
