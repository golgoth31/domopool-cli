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

// filterCmd represents the filter command
var filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		setState, _ := cmd.Flags().GetString("state")
		client := domoClient.NewClient()
		filter := &domopool_proto.Relay{}

		filter.Duration, _ = cmd.Flags().GetUint32("duration")
		filter.State = domopool_proto.RelayStates(domopool_proto.RelayStates_value[setState])
		filter.Relay = domopool_proto.RelayNames(domopool_proto.RelayNames_value["filter"])

		body, _ := proto.Marshal(filter)
		resp := client.Post(fmt.Sprintf("api/%s/%s", viper.GetString("api.version"), "filter"), body)

		if resp.StatusCode() == 200 {
			time.Sleep(2 * time.Second)
			config := domoConfig.GetConfig()

			logger.StdLog.Info().Msgf("%v", config.GetPump())
		}
	},
}

func init() {
	rootCmd.AddCommand(filterCmd)

	filterCmd.Flags().StringP("state", "s", "", "start, stop")
	if err := filterCmd.MarkFlagRequired("state"); err != nil {
		logger.StdLog.Fatal().Err(err).Msg("")
	}
	filterCmd.Flags().Uint32P("duration", "d", 0, "duration of filtering, in minutes")
}
