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

	if response.StatusCode != http.StatusOK {
		// TODO:I revisar todos los tipos
		errorData = ac.ProcessErrorResponse(response)
		return
	}

	account, err = ac.GetAccountFromResponse(response)
	if err != nil {
		errorData = client_error.NewUnknownClientError("Unknown error parsing API GET response")
		return
	}

	return
}
