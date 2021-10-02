package client

import (
	"fmt"
	"github.com/illuque/account-api-client/model"
	"github.com/illuque/account-api-client/model/client_error"
	"net/http"
)

func (ac AccountHttpClient) Fetch(id string) (account *model.AccountData, errorData *client_error.ErrorData) {
	ac.logger.Debugf("Calling API for Fetch with id [%s]...", id)

	getUri := fmt.Sprintf("%s/%s", ac.uri, id)
	response, err := ac.httpClient.Get(getUri)
	if err != nil {
		ac.logger.WithError(err).Errorf("Error sending GET to Account API")
		errorData = client_error.NewUnknownClientError("Unknown error generating API request")
		return
	}

	switch response.StatusCode {
	case http.StatusOK:
		account, errorData = ac.getFetchedAccountFromResponse(response)
	case http.StatusBadRequest:
		errorData = client_error.NewBadRequest("Wrong id parameter format")
	case http.StatusNotFound:
		errorData = client_error.NewNotFound("Specified resource does not exist")
	default:
		errorMsg, _ := ac.getErrorFromResponse(response)
		errorData = client_error.NewUnknownClientError("Unknown error code received from API on GET: " + errorMsg)
	}

	return
}

func (ac AccountHttpClient) getFetchedAccountFromResponse(response *http.Response) (account *model.AccountData, errorData *client_error.ErrorData) {
	account, err := ac.getAccountFromResponse(response)
	if err != nil {
		errorData = client_error.NewUnknownClientError("Unknown error parsing API GET response")
		return
	}
	return
}
