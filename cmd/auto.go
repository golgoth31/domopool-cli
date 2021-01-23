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

// autoCmd represents the filter command
var autoCmd = &cobra.Command{
	Use:   "auto",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		scheme := "http"
		domoClient := resty.New()
		config := &domopool_proto.Config{}

		domoClient.HostURL = scheme + "://192.168.11.183"
		domoClient.SetRetryCount(3)
		domoClient.SetRetryWaitTime(5 * time.Second)

		resp, err := domoClient.
			R().
			Post("/api/v1/auto")
		if err != nil {
			fmt.Println(err)
		}

		if resp.StatusCode() == 200 {
			time.Sleep(2 * time.Second)
			response, err := domoClient.R().Get("/api/v1/config")
			if err != nil {
				fmt.Println(err)
			}
			err = proto.Unmarshal(response.Body(), config)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(config.GetPump())
		}
	},
}

func init() {
	rootCmd.AddCommand(autoCmd)
}