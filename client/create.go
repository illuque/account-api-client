package client

import (
	"bytes"
	"encoding/json"
	"github.com/illuque/account-api-client/model"
	"net/http"
)

type accountCreate struct {
	AccountData model.AccountData `json:"data"`
}

func (ac AccountHttpClient) Create(accountData model.AccountData) (createdAccount *model.AccountData, errorData *model.ErrorData) {
	accountCreateJson, err := json.Marshal(&accountCreate{AccountData: accountData})
	if err != nil {
		ac.logger.WithError(err).Errorf("Error marshaling input account data on Create")
		errorData = model.NewBadRequest("Provided account is not valid, please check AccountData schema")
		return
	}

	ac.logger.Debugf("Calling API for Create with payload [%s]...", accountCreateJson)

	response, err := ac.httpClient.Post(ac.uri, ac.contentType, bytes.NewReader(accountCreateJson))
	if err != nil {
		ac.logger.WithError(err).Errorf("Error sending POST to Account API")
		errorData = model.NewUnknownClientError("Unknown error generating API request")
		return
	}

	switch response.StatusCode {
	case http.StatusCreated:
		createdAccount, err = ac.getAccountFromResponse(response)
		if err != nil {
			errorData = model.NewUnknownClientError("Unknown error parsing API POST response")
			return
		}
	case http.StatusBadRequest:
		errorData = model.NewBadRequest("Wrong parameter(s) provided")
	case http.StatusConflict:
		errorData = model.NewConflict("Specified account already exists")
	default:
		errorMsg, _ := ac.getErrorFromResponse(response)
		errorData = model.NewUnknownClientError("Unknown error code received from API on POST: " + errorMsg)
	}

	return
}
