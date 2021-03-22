package domoConfig

import (
	"fmt"

	"github.com/golgoth31/domopool-cli/internal/domoClient"
	logger "github.com/golgoth31/domopool-cli/internal/log"
	domopool_proto "github.com/golgoth31/domopool-proto"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"
)

func GetConfig() *domopool_proto.Config {
	config := &domopool_proto.Config{}
	client := domoClient.NewClient()
	resp := client.Get(fmt.Sprintf("api/%s/%s", viper.GetString("api.version"), viper.GetString("api.path.config")))

	if err := proto.Unmarshal(resp.Body(), config); err != nil {
		logger.StdLog.Fatal().Err(err).Msg("Unable to unmarchal proto")
	}

	if resp.StatusCode() != 200 {
		logger.StdLog.Fatal().Msg(resp.Status())
	}

	return config
}
