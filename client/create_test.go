package client

import (
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
			case "fails when duplicated account is created":
				// run again to generate duplicate
				// TODO:I meter en un seed en vez de ejecutar a mano!
				gotCreatedAccount, gotErrData = accountHttpClient.Create(tt.args.accountData)
			}

			if (gotErrData != nil) && (gotErrData.Code != tt.wantErrData.Code || gotErrData.Retryable != tt.wantErrData.Retryable) {
				t.Errorf("Create() client_error = %+v, wantErrData %+v", gotErrData, tt.wantErrData)
				return
			}

			if tt.wantId != "" && tt.wantId != gotCreatedAccount.ID { // TODO:I probar de nuevo con el deepEqual
				t.Errorf("Create() gotCreatedAccount = %s, want %s", gotCreatedAccount.ID, tt.wantId)
			}
		})
	}
}

func buildAccountWithoutName() model.AccountData {
	account := buildNewAccount()
	account.Attributes.Name = nil
	return account
}
