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

	switch response.StatusCode {
	case http.StatusCreated:
		createdAccount, err = ac.getAccountFromResponse(response)
		if err != nil {
			errorData = client_error.NewUnknownClientError("Unknown error parsing API POST response")
			return
		}
	case http.StatusBadRequest:
		errorData = client_error.NewBadRequest("Wrong parameter(s) provided")
	case http.StatusConflict:
		errorData = client_error.NewConflict("Specified account already exists")
	default:
		errorMsg, _ := ac.getErrorFromResponse(response)
		errorData = client_error.NewUnknownClientError("Unknown error code received from API on POST: " + errorMsg)
	}

	return
}
