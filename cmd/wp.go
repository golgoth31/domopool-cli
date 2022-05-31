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
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/golgoth31/domopool-cli/internal/domoClient"
	"github.com/golgoth31/domopool-cli/internal/domoConfig"
	logger "github.com/golgoth31/domopool-cli/internal/log"
	domopool_proto "github.com/golgoth31/domopool-proto"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

// wpCmd represents the water pressure command
var wpCmd = &cobra.Command{
	Use:   "wp",
	Short: "Enable/disble or set water pressure component.",
	ValidArgs: []string{
		"enable",
		"disable",
		"set",
	},
	Run: func(cmd *cobra.Command, args []string) {
		wpThreshold, _ := cmd.Flags().GetUint32("threshold")
		wpThresholdAccuracy, _ := cmd.Flags().GetUint32("accuracy")
		wpAdcPin, _ := cmd.Flags().GetUint32("adc-pin")
		wpVmin, _ := cmd.Flags().GetUint32("vmin")
		wpVmax, _ := cmd.Flags().GetUint32("vmax")
		wpAutoCal, _ := cmd.Flags().GetBool("auto-cal")

		client := domoClient.NewClient()

		resp := &resty.Response{}
		if len(args) != 0 {
			switch args[0] {
			case "enable":
				resp = client.Post("/api/v1/wp/enable", nil)
			case "disable":
				resp = client.Post("/api/v1/wp/disable", nil)
			case "set":
				wpbody := &domopool_proto.AnalogSensor{
					AdcPin:            wpAdcPin,
					Threshold:         wpThreshold,
					ThresholdAccuracy: wpThresholdAccuracy,
					Vmin:              wpVmin,
					Vmax:              wpVmax,
					AutoCal:           wpAutoCal,
				}
				body, _ := proto.Marshal(wpbody)
				resp = client.Post("/api/v1/wp/set", body)
			}
			if resp.StatusCode() == 200 {
				time.Sleep(2 * time.Second)
			}
		}

		config := domoConfig.GetConfig()
		logger.StdLog.Info().Msgf("%v", config.Sensors.GetWp())
	},
}

func init() {
	rootCmd.AddCommand(wpCmd)

	wpCmd.Flags().Uint32("accuracy", 8, "set wp threshold accuracy, in %")
	wpCmd.Flags().Uint32("adc-pin", 3, "set wp adc pin")
	wpCmd.Flags().Uint32("vmin", 500000, "set wp threshold accuracy")
	wpCmd.Flags().Uint32("vmax", 4500000, "set wp threshold accuracy")
	wpCmd.Flags().Uint32("threshold", 500000, "set wp sensor threshold")
	wpCmd.Flags().Bool("auto-cal", true, "set wp autocalibration")
}
