package client

import (
	"fmt"
	"github.com/illuque/account-api-client/client/model"
	"net/http"
)

func (ac AccountHttpClient) Delete(id model.DeleteId) (deleted bool, err error) {
	ac.logger.Debugf("Calling Account API for Delete with id [%+v]", id)

	deleteUri := fmt.Sprintf("%s/%s?version=%d", ac.uri, id.Id, id.Version)

	// Create request
	deleteRequest, err := http.NewRequest("DELETE", deleteUri, nil)
	if err != nil {
		ac.logger.WithError(err).Errorf("Error generating DELETE to Account API")
		err = model.NewUnknownError("Unknown error generating API request")
		return
	}

	// Fetch Request
	response, err := ac.httpClient.Do(deleteRequest)
	if err != nil {
		ac.logger.WithError(err).Errorf("Error sending DELETE to Account API")
		err = model.NewUnknownError("Unknown error generating API request")
		return
	}

	switch response.StatusCode {
	case http.StatusNoContent:
		deleted = true
	case http.StatusNotFound:
		err = model.NewNotFound("Specified resource does not exist")
	case http.StatusConflict:
		err = model.NewConflict("Specified version incorrect")
	default:
		errorMsg, _ := ac.getErrorFromResponse(response)
		err = model.NewUnknownError("Unknown error code received from API on DELETE: " + errorMsg)
	}

	return
}
