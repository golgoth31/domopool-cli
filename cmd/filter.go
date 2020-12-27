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
		scheme := "http"
		domoClient := resty.New()
		filter := &domopool_proto.Filter{}

		domoClient.HostURL = scheme + "://192.168.11.183"
		domoClient.SetRetryCount(3)
		domoClient.SetRetryWaitTime(5 * time.Second)

		if setState == "" {
			resp, err := domoClient.R().Get("/api/v1/filter")
			if err != nil {
				fmt.Println(err)
			}
			// err = json.Unmarshal(resp.Body(), config)
			// fmt.Println(resp.String())
			err = proto.Unmarshal(resp.Body(), filter)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(filter)
		} else {
			filter.Duration, _ = cmd.Flags().GetUint32("duration")
			switch setState {
			case "auto":
				filter.State = domopool_proto.FilterStates_auto
			case "start":
				filter.State = domopool_proto.FilterStates_start
			case "stop":
				filter.State = domopool_proto.FilterStates_stop
			}
			body, _ := proto.Marshal(filter)
			resp, err := domoClient.
				R().
				SetBody(body).
				Post("/api/v1/filter")
			if err != nil {
				fmt.Println(err)
			}

			if resp.StatusCode() == 200 {
				time.Sleep(2 * time.Second)
				response, err := domoClient.R().Get("/api/v1/filter")
				if err != nil {
					fmt.Println(err)
				}
				// err = json.Unmarshal(resp.Body(), config)
				// fmt.Println(resp.String())
				err = proto.Unmarshal(response.Body(), filter)
				if err != nil {
					fmt.Println(err)
				}

				fmt.Println(filter)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(filterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// filterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// filterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	filterCmd.Flags().StringP("state", "s", "", "Help message for toggle")
	filterCmd.Flags().Uint32P("duration", "d", 0, "Help message for toggle")
}
