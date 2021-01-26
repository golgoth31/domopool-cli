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
	// Args: cobra.ExactValidArgs(1),
	ValidArgs: []string{
		"enable",
		"disable",
		"set",
	},
	Run: func(cmd *cobra.Command, args []string) {
		wpThreshold, _ := cmd.Flags().GetFloat32("threshold")
		wpThresholdAccuracy, _ := cmd.Flags().GetUint32("accuracy")
		wpAdcPin, _ := cmd.Flags().GetUint32("adc-pin")
		wpVmin, _ := cmd.Flags().GetFloat32("vmin")
		wpVmax, _ := cmd.Flags().GetFloat32("vmax")
		wpAutoCal, _ := cmd.Flags().GetBool("auto-cal")

		scheme := "http"
		domoClient := resty.New()
		wp := &domopool_proto.Config{}

		domoClient.HostURL = scheme + "://192.168.11.183"
		domoClient.SetRetryCount(3)
		domoClient.SetRetryWaitTime(5 * time.Second)

		resp := &resty.Response{}
		var err error
		if len(args) != 0 {
			switch args[0] {
			case "enable":
				resp, err = domoClient.
					R().
					Post("/api/v1/wp/enable")
				if err != nil {
					fmt.Println(err)
				}
			case "disable":
				resp, err = domoClient.
					R().
					Post("/api/v1/wp/disable")
				if err != nil {
					fmt.Println(err)
				}
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
				resp, err = domoClient.
					R().
					SetBody(body).
					Post("/api/v1/wp/set")
				if err != nil {
					fmt.Println(err)
				}
			}
			if resp.StatusCode() == 200 {
				time.Sleep(2 * time.Second)
			}
		}

		readResp, err := domoClient.R().Get("/api/v1/config")
		if err != nil {
			fmt.Println(err)
		}
		err = proto.Unmarshal(readResp.Body(), wp)
		if err != nil {
			fmt.Println(err)
		}
		if readResp.StatusCode() == 200 {
			fmt.Println(wp.Sensors.GetWp())
		}
	},
}

func init() {
	rootCmd.AddCommand(wpCmd)

	wpCmd.Flags().Uint32("accuracy", 8, "set wp threshold accuracy, in %")
	wpCmd.Flags().Uint32("adc-pin", 3, "set wp adc pin")
	wpCmd.Flags().Float32("vmin", 0.5, "set wp threshold accuracy")
	wpCmd.Flags().Float32("vmax", 4.5, "set wp threshold accuracy")
	wpCmd.Flags().Float32("threshold", 0.5, "set wp sensor threshold")
	wpCmd.Flags().Bool("auto-cal", true, "set wp autocalibration")
}
