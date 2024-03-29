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
	domopool_proto "github.com/golgoth31/domopool-proto"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

// lightCmd represents the filter command
var lightCmd = &cobra.Command{
	Use:   "light",
	Short: "The command to turn on or off the pool lights.",
	Run: func(cmd *cobra.Command, args []string) {
		setState, _ := cmd.Flags().GetString("state")
		scheme := "http"
		domoClient := resty.New()
		relay := &domopool_proto.Relay{}

		domoClient.HostURL = scheme + "://192.168.11.183"
		domoClient.SetRetryCount(3)
		domoClient.SetRetryWaitTime(5 * time.Second)

		relay.Duration = 0
		relay.State = domopool_proto.RelayStates(domopool_proto.RelayStates_value[setState])
		relay.Relay = domopool_proto.RelayNames(domopool_proto.RelayNames_value["light"])
		body, _ := proto.Marshal(relay)
		resp, err := domoClient.
			R().
			SetBody(body).
			Post("/api/v1/light")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp.Status())
	},
}

func init() {
	rootCmd.AddCommand(lightCmd)

	lightCmd.Flags().StringP("state", "s", "", "Help message for toggle")
	if err := lightCmd.MarkFlagRequired("state"); err != nil {
		fmt.Println("Unable to mark flag as required.")
	}
}
