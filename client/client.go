package client

import (
	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/illuque/account-api-client/client/model"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

type AccountClient interface {
	Create(account model.AccountData) (createdAccount *model.AccountData, err error)
	Fetch(id string) (account *model.AccountData, err error)
	Delete(id model.DeleteId) (deleted bool, err error)
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
		},
		uri:         uri,
		contentType: "application/vnd.api+json",
		logger:      logger,
	}
}
