package client

import (
	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/illuque/account-api-client/model"
	"github.com/illuque/account-api-client/model/client_error"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

// TODO:I Test this

type AccountClient interface {
	Create(account model.AccountData) (createdAccount *model.AccountData, errorData *client_error.ErrorData)
	Fetch(id string) (account *model.AccountData, errorData *client_error.ErrorData)
	Delete(id string, version int64) (deleted bool, errorData *client_error.ErrorData)
}

type AccountHttpClient struct {
	httpClient  *http.Client
	uri         string
	contentType string
	logger      *logrus.Logger
}

func NewAccountApiClient(uri string, timeout time.Duration) AccountClient {
	logger := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
	}

	formatter := runtime.Formatter{ChildFormatter: &log.JSONFormatter{}}
	formatter.Line = true
	formatter.File = true
	logger.SetFormatter(&formatter)

	return AccountHttpClient{
		httpClient: &http.Client{
			Timeout: timeout,
		}, // TODO:I ver si hay una manera mejor de inicializarlo
		uri:         uri,
		contentType: "application/vnd.api+json",
		logger:      logger,
	}
}
