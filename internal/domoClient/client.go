package domoClient

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	logger "github.com/golgoth31/domopool-cli/internal/log"
	"github.com/spf13/viper"
)

func NewClient() *DomoClient {
	var client DomoClient
	client.Client = resty.New()

	client.Client.HostURL = fmt.Sprintf(
		"%s://%s:%d",
		viper.GetString("boxScheme"),
		viper.GetString("boxHost"),
		viper.GetInt("boxPort"),
	)
	client.Client.SetRetryCount(3)
	client.Client.SetRetryWaitTime(5 * time.Second)

	return &client
}

func (c *DomoClient) Get(path string) *resty.Response {
	resp, err := c.Client.R().Get(path)
	if err != nil {
		logger.StdLog.Fatal().Err(err).Msg("Unable to request domopool box")
	}

	return resp
}

func (c *DomoClient) Post(path string, body interface{}) *resty.Response {
	resp, err := c.
		Client.
		R().
		SetBody(body).
		Post(path)
	if err != nil {
		logger.StdLog.Fatal().Err(err).Msg("Unable to request domopool box")
	}

	return resp
}
