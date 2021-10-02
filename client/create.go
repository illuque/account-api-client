package client

import (
	"bytes"
	"encoding/json"
	"github.com/illuque/account-api-client/model"
	"github.com/illuque/account-api-client/model/client_error"
	"net/http"
)

type accountCreate struct { // TODO:I ver si mover a common
	AccountData model.AccountData `json:"data"`
}

// TODO:I leer si es mejor usar http.AccountClient{} o http.Post, etc.

func (ac AccountHttpClient) Create(accountData model.AccountData) (createdAccount *model.AccountData, errorData *client_error.ErrorData) {
	accountCreateJson, err := json.Marshal(&accountCreate{AccountData: accountData})
	if err != nil {
		ac.logger.WithError(err).Errorf("Error marshaling input account data on Create")
		errorData = client_error.NewBadRequest("Provided account is not valid, please check AccountData schema")
		return
	}

	ac.logger.Debugf("Calling API for Create with payload [%s]...", accountCreateJson)

	response, err := ac.httpClient.Post(ac.uri, ac.contentType, bytes.NewReader(accountCreateJson))
	if err != nil {
		ac.logger.WithError(err).Errorf("Error sending POST to Account API")
		errorData = client_error.NewUnknownClientError("Unknown error generating API request")
		return
	}

	if response.StatusCode != http.StatusCreated {
		errorData = ac.ProcessErrorResponse(response)
		return
	}

	createdAccount, err = ac.GetAccountFromResponse(response)
	if err != nil {
		errorData = client_error.NewUnknownClientError("Unknown error parsing API POST response")
		return
	}

	return
}
