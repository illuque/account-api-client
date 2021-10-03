package client

import (
	"github.com/illuque/account-api-client/client/model"
	"reflect"
	"testing"
)

func TestAccountHttpClient_Create(t *testing.T) {
	accountHttpClient := buildClient()

	type args struct {
		accountData model.AccountData
	}

	account := buildNewAccount()

	tests := []struct {
		name               string
		args               args
		wantCreatedAccount *model.AccountData
		wantAccount        *model.AccountData
		wantErr            error
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
			wantErr: model.NewConflict("Specified account already exists"),
		},
		{
			name: "bad request when name not provided",
			args: args{
				buildAccountWithoutName(),
			},
			wantErr: model.NewBadRequest("Wrong parameter(s) provided"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			switch tt.name {
			case "conflict when duplicated account is created":
				// run again to generate duplicate
				_, err := accountHttpClient.Create(tt.args.accountData)
				if err != nil {
					t.FailNow()
				}
			}

			gotCreatedAccount, gotErr := accountHttpClient.Create(tt.args.accountData)

			if !reflect.DeepEqual(gotCreatedAccount, tt.wantAccount) {
				t.Errorf("Create() gotCreatedAccount = %v, want %v", gotCreatedAccount, tt.wantAccount)
			}

			if (gotErr != nil) && !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Create() client_error = %+v, wantErr %+v", gotErr, tt.wantErr)
				return
			}
		})
	}
}

func buildAccountWithoutName() model.AccountData {
	account := buildNewAccount()
	account.Attributes.Name = nil
	return account
}
