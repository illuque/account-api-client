package client

import (
	"fmt"
	model2 "github.com/illuque/account-api-client/client/model"
	"net/http"
)

func (ac AccountHttpClient) Fetch(id string) (account *model2.AccountData, errorData *model2.ErrorData) {
	ac.logger.Debugf("Calling API for Fetch with id [%s]...", id)

	getUri := fmt.Sprintf("%s/%s", ac.uri, id)
	response, err := ac.httpClient.Get(getUri)
	if err != nil {
		ac.logger.WithError(err).Errorf("Error sending GET to Account API")
		errorData = model2.NewUnknownClientError("Unknown error generating API request")
		return
	}

	switch response.StatusCode {
	case http.StatusOK:
		account, errorData = ac.getFetchedAccountFromResponse(response)
	case http.StatusBadRequest:
		errorData = model2.NewBadRequest("Wrong id parameter format")
	case http.StatusNotFound:
		errorData = model2.NewNotFound("Specified resource does not exist")
	default:
		errorMsg, _ := ac.getErrorFromResponse(response)
		errorData = model2.NewUnknownClientError("Unknown error code received from API on GET: " + errorMsg)
	}

	return
}

func (ac AccountHttpClient) getFetchedAccountFromResponse(response *http.Response) (account *model2.AccountData, errorData *model2.ErrorData) {
	account, err := ac.getAccountFromResponse(response)
	if err != nil {
		errorData = model2.NewUnknownClientError("Unknown error parsing API GET response")
		return
	}
	return
}
