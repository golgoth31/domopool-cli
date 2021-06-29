/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/golgoth31/domopool-cli/internal/domoClient"
	"github.com/golgoth31/domopool-cli/internal/domoConfig"
	logger "github.com/golgoth31/domopool-cli/internal/log"
	domopool_proto "github.com/golgoth31/domopool-proto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"
)

// relayCmd represents the filter command
var relayCmd = &cobra.Command{
	Use:   "relay",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		setState, _ := cmd.Flags().GetString("state")
		setCh, _ := cmd.Flags().GetBool("ch")
		setPh, _ := cmd.Flags().GetBool("ph")
		relay := &domopool_proto.Relay{}

		relay.Duration, _ = cmd.Flags().GetUint32("duration")
		relay.State = domopool_proto.RelayStates(domopool_proto.RelayStates_value[setState])

		if setCh {
			logger.StdLog.Info().Msgf("Setting ch relay: %s", setState)
			relay.Relay = domopool_proto.RelayNames(domopool_proto.RelayNames_value["ch"])

			makeRequest(relay)
			time.Sleep(2 * time.Second)
		}
		if setPh {
			logger.StdLog.Info().Msgf("Setting ch relay: %s", setState)
			relay.Relay = domopool_proto.RelayNames(domopool_proto.RelayNames_value["ph"])

			makeRequest(relay)
		}

		if !setCh && !setPh {
			logger.StdLog.Info().Msgf("Setting filter relay: %s", setState)
			relay.Relay = domopool_proto.RelayNames(domopool_proto.RelayNames_value["filter"])

			makeRequest(relay)
		}
	},
}

func makeRequest(relay *domopool_proto.Relay) {
	client := domoClient.NewClient()
	body, _ := proto.Marshal(relay)
	resp := client.Post(fmt.Sprintf("api/%s/%s", viper.GetString("api.version"), "relay"), body)

	if resp.StatusCode() == 200 {
		time.Sleep(2 * time.Second)
		config := domoConfig.GetConfig()

		logger.StdLog.Info().Msgf("%v", config.GetPump())
	}
}

func init() {
	rootCmd.AddCommand(relayCmd)

	relayCmd.Flags().StringP("state", "s", "", "start, stop")
	relayCmd.Flags().Bool("ch", false, "Ch relay")
	relayCmd.Flags().Bool("ph", false, "Ph relay")
	relayCmd.Flags().Uint32P("duration", "d", 0, "duration of filtering, in minutes")

	if err := relayCmd.MarkFlagRequired("state"); err != nil {
		logger.StdLog.Fatal().Err(err).Msg("")
	}
}
