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
	"github.com/golgoth31/domopool-cli/internal/domoClient"
	logger "github.com/golgoth31/domopool-cli/internal/log"
	domopool_proto "github.com/golgoth31/domopool-proto"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

// adcCmd represents the filter command
var adcCmd = &cobra.Command{
	Use:   "adc",
	Short: "Configure the adc for anolog measurements.",
	Run: func(cmd *cobra.Command, args []string) {
		ADCMode, _ := cmd.Flags().GetUint32("mode")
		ADCGain, _ := cmd.Flags().GetUint32("gain")
		ADCDatarate, _ := cmd.Flags().GetUint32("datarate")
		adc := &domopool_proto.Sensors{}
		client := domoClient.NewClient()
		adc.AdcDatarate = ADCDatarate
		adc.AdcGain = ADCGain
		adc.AdcMode = ADCMode
		body, _ := proto.Marshal(adc)
		resp := client.Post("/api/v1/adc/set", body)

		logger.StdLog.Info().Msg(resp.Status())
	},
}

func init() {
	rootCmd.AddCommand(adcCmd)

	adcCmd.Flags().Uint32P("mode", "m", 1, "ADC mode (1: single; 0: continuous)")
	adcCmd.Flags().Uint32P("gain", "g", 1, "ADC gain")
	adcCmd.Flags().Uint32P("datarate", "d", 1, "ADC datarate")
}
