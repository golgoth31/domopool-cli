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

	"github.com/go-resty/resty/v2"
	"github.com/gogo/protobuf/proto"
	domopool_proto "github.com/golgoth31/domopool-proto"
	"github.com/spf13/cobra"
)

// adcCmd represents the filter command
var adcCmd = &cobra.Command{
	Use:   "adc",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactValidArgs(1),
	ValidArgs: []string{
		"set",
	},
	Run: func(cmd *cobra.Command, args []string) {
		ADCMode, _ := cmd.Flags().GetUint32("mode")
		ADCGain, _ := cmd.Flags().GetUint32("gain")
		ADCDatarate, _ := cmd.Flags().GetUint32("datarate")

		scheme := "http"
		domoClient := resty.New()
		adc := &domopool_proto.Sensors{}

		domoClient.HostURL = scheme + "://192.168.11.183"
		domoClient.SetRetryCount(3)
		domoClient.SetRetryWaitTime(5 * time.Second)

		adc.AdcDatarate = ADCDatarate
		adc.AdcGain = ADCGain
		adc.AdcMode = ADCMode
		body, _ := proto.Marshal(adc)
		resp, err := domoClient.
			R().
			SetBody(body).
			Post("/api/v1/adc/set")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp.Status())
	},
}

func init() {
	rootCmd.AddCommand(adcCmd)

	adcCmd.Flags().Uint32P("mode", "m", 1, "ADC mode (1: single; 0: continuous)")
	adcCmd.Flags().Uint32P("gain", "g", 1, "ADC gain")
	adcCmd.Flags().Uint32P("datarate", "d", 4, "ADC datarate")
}
