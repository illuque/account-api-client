package client

import (
	"github.com/google/uuid"
	"github.com/illuque/account-api-client/model"
	"github.com/illuque/account-api-client/model/client_error"
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
		wantAccount   *model.AccountData
		wantErrorData *client_error.ErrorData
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
			wantErrorData: client_error.NewFromApiError(404, "whatever"),
		}, {
			name: "bad request when invalid id",
			args: args{
				id: "fake id",
			},
			wantAccount:   nil,
			wantErrorData: client_error.NewFromApiError(400, "whatever"),
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
			if (gotErrorData != nil) && (gotErrorData.Code != tt.wantErrorData.Code || gotErrorData.Retryable != tt.wantErrorData.Retryable) {
				t.Errorf("Fetch() gotErrorData = %v, want %v", gotErrorData, tt.wantErrorData)
				return
			}
		})
	}
}
