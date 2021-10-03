package client

import (
	"bytes"
	"encoding/json"
	model2 "github.com/illuque/account-api-client/client/model"
	"net/http"
)

type accountCreate struct {
	AccountData model2.AccountData `json:"data"`
}

func (ac AccountHttpClient) Create(accountData model2.AccountData) (createdAccount *model2.AccountData, errorData *model2.ErrorData) {
	accountCreateJson, err := json.Marshal(&accountCreate{AccountData: accountData})
	if err != nil {
		ac.logger.WithError(err).Errorf("Error marshaling input account data on Create")
		errorData = model2.NewBadRequest("Provided account is not valid, please check AccountData schema")
		return
	}

	ac.logger.Debugf("Calling API for Create with payload [%s]...", accountCreateJson)

	response, err := ac.httpClient.Post(ac.uri, ac.contentType, bytes.NewReader(accountCreateJson))
	if err != nil {
		ac.logger.WithError(err).Errorf("Error sending POST to Account API")
		errorData = model2.NewUnknownClientError("Unknown error generating API request")
		return
	}

	switch response.StatusCode {
	case http.StatusCreated:
		createdAccount, err = ac.getAccountFromResponse(response)
		if err != nil {
			errorData = model2.NewUnknownClientError("Unknown error parsing API POST response")
			return
		}
	case http.StatusBadRequest:
		errorData = model2.NewBadRequest("Wrong parameter(s) provided")
	case http.StatusConflict:
		errorData = model2.NewConflict("Specified account already exists")
	default:
		errorMsg, _ := ac.getErrorFromResponse(response)
		errorData = model2.NewUnknownClientError("Unknown error code received from API on POST: " + errorMsg)
	}

	return
}
