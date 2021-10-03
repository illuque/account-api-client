package client

import (
	"github.com/google/uuid"
	model2 "github.com/illuque/account-api-client/client/model"
	"reflect"
	"testing"
	"time"
)

func TestAccountHttpClient_Fetch(t *testing.T) {
	accountHttpClient := NewAccountApiClient("http://localhost:8080/v1/organisation/accounts", 2*time.Second)

	type args struct {
		id string
	}

	account := buildNewAccount()

	tests := []struct {
		name          string
		args          args
		wantAccount   *model2.AccountData
		wantErrorData *model2.ErrorData
	}{
		{
			name: "retrieved when existing",
			args: args{
				id: account.ID,
			},
			wantAccount:   &account,
			wantErrorData: nil,
		}, {
			name: "not found when not existing",
			args: args{
				id: uuid.New().String(),
			},
			wantAccount:   nil,
			wantErrorData: model2.NewNotFound("Specified resource does not exist"),
		}, {
			name: "bad request when invalid id",
			args: args{
				id: "fake id",
			},
			wantAccount:   nil,
			wantErrorData: model2.NewBadRequest("Wrong id parameter format"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			switch tt.name {
			case "retrieved when existing":
				// TODO:I meter en un seed en vez de ejecutar a mano!
				_, err := accountHttpClient.Create(account)
				if err != nil {
					t.FailNow()
				}
			}

			gotAccount, gotErrorData := accountHttpClient.Fetch(tt.args.id)

			if !reflect.DeepEqual(gotAccount, tt.wantAccount) {
				t.Errorf("Fetch() gotAccount = %v, want %v", gotAccount, tt.wantAccount)
			}

			if !reflect.DeepEqual(gotErrorData, tt.wantErrorData) {
				t.Errorf("Fetch() gotErrorData = %v, want %v", gotErrorData, tt.wantErrorData)
				return
			}
		})
	}
}
