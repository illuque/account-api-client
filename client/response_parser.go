package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/illuque/account-api-client/model"
	"github.com/illuque/account-api-client/model/client_error"
	"io/ioutil"
	"net/http"
)

// TODO:I remove
func (ac AccountHttpClient) ProcessErrorResponse(response *http.Response) *client_error.ErrorData {
	apiErrMsg, _ := ac.getErrorFromResponse(response)
	errMsg := fmt.Sprintf("API error with code '%d', message '%s'", response.StatusCode, apiErrMsg)
	err := errors.New(errMsg)
	ac.logger.WithError(err).Errorf("API responded '%d' on Create", response.StatusCode)

	return client_error.NewFromApiError(response.StatusCode, apiErrMsg)
}

func (ac AccountHttpClient) GetAccountFromResponse(response *http.Response) (responseAccount *model.AccountData, parseError error) {
	defer response.Body.Close()

	bodyBytes, parseError := ioutil.ReadAll(response.Body)
	if parseError != nil {
		ac.logger.WithError(parseError).Errorf("Error reading API response")
		return
	}

	var accountResponseParsed accountCreate
	if parseError = json.Unmarshal(bodyBytes, &accountResponseParsed); parseError != nil {
		ac.logger.WithError(parseError).Errorf("Error reading error response as AccountData")
		return
	}

	responseAccount = &accountResponseParsed.AccountData

	return
}

func (ac AccountHttpClient) getErrorFromResponse(response *http.Response) (errorMessage string, parseError error) {
	defer response.Body.Close()

	bodyBytes, parseError := ioutil.ReadAll(response.Body)
	if parseError != nil {
		ac.logger.WithError(parseError).Errorf("Error reading API response")
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
