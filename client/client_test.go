package client

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestNewAccountApiClient(t *testing.T) {
	type args struct {
		uri     string
		timeout time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "all ok",
			args: args{
				uri:     "https://test.com",
				timeout: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewAccountApiClient(tt.args.uri, tt.args.timeout)
			assert.NotNil(t, got)

			gotReflect := reflect.Indirect(reflect.ValueOf(got))

			uri := gotReflect.FieldByName("uri")
			assert.Equal(t, uri.String(), tt.args.uri)

			contentType := gotReflect.FieldByName("contentType")
			assert.Equal(t, contentType.String(), "application/vnd.api+json")

			httpClientReflect := reflect.Indirect(gotReflect.FieldByName("httpClient"))
			timeout := httpClientReflect.FieldByName("Timeout")
			assert.NotNil(t, timeout.Int(), tt.args.timeout)
		})
	}
}
