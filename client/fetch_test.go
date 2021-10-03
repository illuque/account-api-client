package client

import (
	"github.com/google/uuid"
	"github.com/illuque/account-api-client/client/model"
	"reflect"
	"testing"
)

func TestAccountHttpClient_Fetch(t *testing.T) {
	accountHttpClient := buildClient()

	type args struct {
		id string
	}

	account := buildNewAccount()

	tests := []struct {
		name        string
		args        args
		wantAccount *model.AccountData
		wantErr     error
	}{
		{
			name: "retrieved when existing",
			args: args{
				id: account.ID,
			},
			wantAccount: &account,
			wantErr:     nil,
		}, {
			name: "not found when not existing",
			args: args{
				id: uuid.New().String(),
			},
			wantAccount: nil,
			wantErr:     model.NewNotFound("Specified resource does not exist"),
		}, {
			name: "bad request when invalid id",
			args: args{
				id: "fake id",
			},
			wantAccount: nil,
			wantErr:     model.NewBadRequest("Wrong id parameter format"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			switch tt.name {
			case "retrieved when existing":
				// create one to ensure existing
				_, err := accountHttpClient.Create(account)
				if err != nil {
					t.FailNow()
				}
			}

			gotAccount, gotErr := accountHttpClient.Fetch(tt.args.id)

			if !reflect.DeepEqual(gotAccount, tt.wantAccount) {
				t.Errorf("Fetch() gotAccount = %v, want %v", gotAccount, tt.wantAccount)
			}

			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Fetch() gotErr = %v, want %v", gotErr, tt.wantErr)
				return
			}
		})
	}
}
