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
	Create(account model.AccountData) (*model.AccountData, *client_error.ErrorData)
	Fetch(id string) (*model.AccountData, *client_error.ErrorData)
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
	} // TODO:I pintar timestamps

	formatter := runtime.Formatter{ChildFormatter: &log.TextFormatter{}}
	logger.SetFormatter(&formatter)
	formatter.Line = true
	formatter.File = true

	return AccountHttpClient{
		httpClient: &http.Client{
			Timeout: timeout,
		}, // TODO:I ver si hay una manera mejor de inicializarlo
		uri:         uri,
		contentType: "application/vnd.api+json",
		logger:      logger,
	}
}
