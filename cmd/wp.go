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

// wpCmd represents the filter command
var wpCmd = &cobra.Command{
	Use:   "wp",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactValidArgs(1),
	ValidArgs: []string{
		"enable",
		"disable",
		"set",
	},
	Run: func(cmd *cobra.Command, args []string) {
		wpThreshold, _ := cmd.Flags().GetFloat32("wp")
		wpThresholdAccuracy, _ := cmd.Flags().GetUint32("wp-accuracy")
		wpAdcPin, _ := cmd.Flags().GetUint32("wp-adc-pin")
		wpVmin, _ := cmd.Flags().GetFloat32("wp-vmin")
		wpVmax, _ := cmd.Flags().GetFloat32("wp-vmax")

		scheme := "http"
		domoClient := resty.New()
		wp := &domopool_proto.AnalogSensor{}

		domoClient.HostURL = scheme + "://192.168.11.183"
		domoClient.SetRetryCount(3)
		domoClient.SetRetryWaitTime(5 * time.Second)

		wp.AdcPin = wpAdcPin
		wp.Threshold = wpThreshold
		wp.ThresholdAccuracy = wpThresholdAccuracy
		wp.Vmax = wpVmax
		wp.Vmin = wpVmin
		body, _ := proto.Marshal(wp)
		resp, err := domoClient.
			R().
			SetBody(body).
			Post("/api/v1/wp")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp.Status())
	},
}

func init() {
	rootCmd.AddCommand(wpCmd)

	wpCmd.Flags().Uint32("accuracy", 8, "post wp threshold accuracy")
	wpCmd.Flags().Uint32("adc-pin", 3, "post wp adc pin")
	wpCmd.Flags().Float32("vmin", 0.5, "post wp threshold accuracy")
	wpCmd.Flags().Float32("vmax", 4.5, "post wp threshold accuracy")
	wpCmd.Flags().Bool("threshold", false, "get current water pressure sensor threshold")
}