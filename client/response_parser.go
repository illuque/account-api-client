package client

import (
	"encoding/json"
	"github.com/illuque/account-api-client/model"
	"io/ioutil"
	"net/http"
)

func (ac AccountHttpClient) ParseResponseBodyOnSuccess(response *http.Response) (responseAccount *model.AccountData, parseError error) {
	defer response.Body.Close()

	bodyBytes, parseError := ioutil.ReadAll(response.Body)
	if parseError != nil {
		ac.logger.WithError(parseError).Errorf("Error reading POST response")
		return
	}

	var accountResponseParsed accountCreate
	if parseError = json.Unmarshal(bodyBytes, &accountResponseParsed); parseError != nil {
		ac.logger.WithError(parseError).Errorf("Error unmarshaling POST response to AccountData")
		return
	}

	responseAccount = &accountResponseParsed.AccountData

	return
}

func (ac AccountHttpClient) ParseResponseBodyOnError(response *http.Response) (errorMessage string, parseError error) {
	defer response.Body.Close()

	bodyBytes, parseError := ioutil.ReadAll(response.Body)
	if parseError != nil {
		ac.logger.WithError(parseError).Errorf("Error reading POST response")
		return
	}

	var apiErrorBody model.ApiErrorBody
	if parseError = json.Unmarshal(bodyBytes, &apiErrorBody); parseError != nil {
		ac.logger.WithError(parseError).Errorf("Error reading error response as ApiErrorBody")
		return
	}

	errorMessage = apiErrorBody.Message

	return
}
