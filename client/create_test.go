package client

import (
	"github.com/google/uuid"
	"github.com/illuque/account-api-client/model"
	"github.com/illuque/account-api-client/model/client_error"
	"testing"
	"time"
)

func TestAccountHttpClient_Create(t *testing.T) {
	accountHttpClient := NewAccountApiClient("http://localhost:8080/v1/organisation/accounts", 2*time.Second)

	type args struct {
		accountData model.AccountData
	}

	accountForOk := buildNewAccount()

	tests := []struct {
		name               string
		args               args
		wantCreatedAccount *model.AccountData
		wantId             string
		wantErrData        *client_error.ErrorData
	}{
		{
			name: "succeeds when valid payload",
			args: args{
				accountForOk,
			},
			wantId: accountForOk.ID,
		},
		{
			name: "fails when duplicated account is created",
			args: args{
				buildNewAccount(),
			},
			wantErrData: client_error.NewFromApiError(409, "Account cannot be created as it violates a duplicate constraint"),
		},
		{
			name: "fails when name not provided",
			args: args{
				buildAccountWithoutName(),
			},
			wantErrData: client_error.NewBadRequest("name in body is required"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotCreatedAccount, gotErrData := accountHttpClient.Create(tt.args.accountData)

			switch tt.name {
			case "succeeds when valid payload":
			case "fails when name not provided":
			case "fails when duplicated account is created":
				// run again to generate duplicate
				gotCreatedAccount, gotErrData = accountHttpClient.Create(tt.args.accountData)
			}

			if (gotErrData != nil) && (gotErrData.Code != tt.wantErrData.Code || gotErrData.Retryable != tt.wantErrData.Retryable) {
				t.Errorf("Create() client_error = %+v, wantErrData %+v", gotErrData, tt.wantErrData)
				return
			}

			if tt.wantId != "" && tt.wantId != gotCreatedAccount.ID {
				t.Errorf("Create() gotCreatedAccount = %s, want %s", gotCreatedAccount.ID, tt.wantId)
			}
		})
	}
}

func buildNewAccount() model.AccountData {
	var countryUK = "UK"
	var version = int64(0)

	return model.AccountData{
		Attributes: &model.AccountAttributes{
			BankID:     "400305",
			BankIDCode: "GBDSC",
			Bic:        "LHVBEE22",
			Country:    &countryUK,
			Name: []string{
				"James Bond",
			},
		},
		ID:             uuid.New().String(),
		OrganisationID: "15a63614-6ae1-4f5b-8f43-7d4dfcb37e76",
		Type:           "accounts",
		Version:        &version,
	}
}

func buildAccountWithoutName() model.AccountData {
	account := buildNewAccount()
	account.Attributes.Name = nil
	return account
}
