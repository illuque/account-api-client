package client

import (
	"bytes"
	"encoding/json"
	"github.com/illuque/account-api-client/client/model"
	"net/http"
)

type accountCreate struct {
	AccountData model.AccountData `json:"data"`
}

func (ac AccountHttpClient) Create(accountData model.AccountData) (createdAccount *model.AccountData, err error) {
	accountCreateJson, err := json.Marshal(&accountCreate{AccountData: accountData})
	if err != nil {
		ac.logger.WithError(err).Errorf("Error marshaling input account data on Create")
		err = model.NewBadRequest("Provided account is not valid, please check AccountData schema")
		return
	}

	ac.logger.Debugf("Calling Account API for Create with payload [%s]...", accountCreateJson)

	response, err := ac.httpClient.Post(ac.uri, ac.contentType, bytes.NewReader(accountCreateJson))
	if err != nil {
		ac.logger.WithError(err).Errorf("Error sending POST to Account API")
		err = model.NewUnknownError("Unknown error generating API request")
		return
	}

	switch response.StatusCode {
	case http.StatusCreated:
		createdAccount, err = ac.getAccountFromResponse(response)
		if err != nil {
			err = model.NewUnknownError("Unknown error parsing API POST response")
			return
		}
	case http.StatusBadRequest:
		err = model.NewBadRequest("Wrong parameter(s) provided")
	case http.StatusConflict:
		err = model.NewConflict("Specified account already exists")
	default:
		errorMsg, _ := ac.getErrorFromResponse(response)
		err = model.NewUnknownError("Unknown error code received from API on POST: " + errorMsg)
	}

	return
}
