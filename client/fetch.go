package client

import (
	"fmt"
	"github.com/illuque/account-api-client/client/model"
	"net/http"
)

func (ac AccountHttpClient) Fetch(id string) (account *model.AccountData, err error) {
	ac.logger.Debugf("Calling Account API for Fetch with id [%s]...", id)

	getUri := fmt.Sprintf("%s/%s", ac.uri, id)
	response, err := ac.httpClient.Get(getUri)
	if err != nil {
		ac.logger.WithError(err).Errorf("Error sending GET to Account API")
		err = model.NewUnknownError("Unknown error generating API request")
		return
	}

	switch response.StatusCode {
	case http.StatusOK:
		account, err = ac.getFetchedAccountFromResponse(response)
	case http.StatusBadRequest:
		err = model.NewBadRequest("Wrong id parameter format")
	case http.StatusNotFound:
		err = model.NewNotFound("Specified resource does not exist")
	default:
		errorMsg, _ := ac.getErrorFromResponse(response)
		err = model.NewUnknownError("Unknown error code received from API on GET: " + errorMsg)
	}

	return
}

func (ac AccountHttpClient) getFetchedAccountFromResponse(response *http.Response) (account *model.AccountData, err error) {
	account, errApi := ac.getAccountFromResponse(response)
	if errApi != nil {
		err = model.NewUnknownError("Unknown error parsing API GET response")
		return
	}
	return
}
