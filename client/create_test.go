package client

import (
	model2 "github.com/illuque/account-api-client/client/model"
	"reflect"
	"testing"
	"time"
)

func TestAccountHttpClient_Create(t *testing.T) {
	accountHttpClient := NewAccountApiClient("http://localhost:8080/v1/organisation/accounts", 2*time.Second)

	type args struct {
		accountData model2.AccountData
	}

	account := buildNewAccount()

	tests := []struct {
		name               string
		args               args
		wantCreatedAccount *model2.AccountData
		wantAccount        *model2.AccountData
		wantErrorData      *model2.ErrorData
	}{
		{
			name: "succeeds when valid payload and first version",
			args: args{
				account,
			},
			wantAccount: &account,
		},
		{
			name: "conflict when duplicated account is created",
			args: args{
				buildNewAccount(),
			},
			wantErrorData: model2.NewConflict("Specified account already exists"),
		},
		{
			name: "bad request when name not provided",
			args: args{
				buildAccountWithoutName(),
			},
			wantErrorData: model2.NewBadRequest("Wrong parameter(s) provided"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotCreatedAccount, gotErrorData := accountHttpClient.Create(tt.args.accountData)

			switch tt.name {
			case "conflict when duplicated account is created":
				// run again to generate duplicate
				// TODO:I meter en un seed en vez de ejecutar a mano!
				gotCreatedAccount, gotErrorData = accountHttpClient.Create(tt.args.accountData)
			}

			if !reflect.DeepEqual(gotCreatedAccount, tt.wantAccount) {
				t.Errorf("Create() gotCreatedAccount = %v, want %v", gotCreatedAccount, tt.wantAccount)
			}

			if (gotErrorData != nil) && !reflect.DeepEqual(gotErrorData, tt.wantErrorData) {
				t.Errorf("Create() client_error = %+v, wantErrorData %+v", gotErrorData, tt.wantErrorData)
				return
			}
		})
	}
}

func buildAccountWithoutName() model2.AccountData {
	account := buildNewAccount()
	account.Attributes.Name = nil
	return account
}
