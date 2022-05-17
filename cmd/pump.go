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

// pumpCmd represents the filter command
var pumpCmd = &cobra.Command{
	Use:   "pump",
	Short: "Change pump timing (dynqmic, half day, full day)",
	Run: func(cmd *cobra.Command, args []string) {
		setTime, _ := cmd.Flags().GetString("time")
		pump := &domopool_proto.Pump{}
		pump.Timing = domopool_proto.PumpTiming(domopool_proto.PumpTiming_value[setTime])

		makePumpRequest(pump)
	},
}

func makePumpRequest(pump *domopool_proto.Pump) {
	client := domoClient.NewClient()
	body, _ := proto.Marshal(pump)
	resp := client.Post(fmt.Sprintf("api/%s/%s", viper.GetString("api.version"), "pump/time"), body)

	if resp.StatusCode() == 200 {
		time.Sleep(2 * time.Second)
		config := domoConfig.GetConfig()

		logger.StdLog.Info().Msgf("%v", config.GetPump())
	}
}

func init() {
	rootCmd.AddCommand(pumpCmd)

	pumpCmd.Flags().String("time", "dynamic", "Pump timing (dynamic, half_day, full_day)")

	if err := pumpCmd.MarkFlagRequired("time"); err != nil {
		logger.StdLog.Fatal().Err(err).Msg("")
	}
}
