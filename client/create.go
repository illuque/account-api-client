package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/illuque/account-api-client/model"
	"github.com/illuque/account-api-client/model/client_error"
	"net/http"
)

type accountCreate struct {
	AccountData model.AccountData `json:"data"`
}

// TODO:I leer si es mejor usar http.AccountClient{} o http.Post, etc.

func (ac AccountHttpClient) Create(accountData model.AccountData) (createdAccount *model.AccountData, errorData *client_error.ErrorData) {
	ac.logger.Info("Calling API for Create...")

	accountCreateJson, err := json.Marshal(&accountCreate{AccountData: accountData})
	if err != nil {
		ac.logger.WithError(err).Errorf("Error marshaling input account data on Create")
		errorData = client_error.NewBadRequest("Provided account is not valid, please check AccountData schema")
		return
	}

	ac.logger.Debugf("Payload to send [%s]", accountCreateJson)

	response, err := ac.client.Post(ac.uri, ac.contentType, bytes.NewReader(accountCreateJson))
	if err != nil {
		ac.logger.WithError(err).Errorf("Error on account POST")
		errorData = client_error.NewUnknownClientError("Unknown error generating API request")
		return
	}

	if response.StatusCode != http.StatusCreated {
		apiErrMsg, _ := ac.ParseResponseBodyOnError(response)
		errMsg := fmt.Sprintf("client_error status code '%d', message '%s'", response.StatusCode, apiErrMsg)
		err = errors.New(errMsg)

		errorData = client_error.NewFromApiError(response.StatusCode, apiErrMsg)

		ac.logger.WithError(err).Errorf("API responded '%d' on Create", response.StatusCode)
		return
	}

	createdAccount, err = ac.ParseResponseBodyOnSuccess(response)
	if err != nil {
		errorData = client_error.NewUnknownClientError("Unknown error parsing API response")
		return
	}

	return
}
