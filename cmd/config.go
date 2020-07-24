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
	"encoding/json"
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	domopool_proto "github.com/golgoth31/domopool-proto"
	"github.com/spf13/cobra"
	"gopkg.in/resty.v1"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getConfig, _ := cmd.Flags().GetString("get")
		scheme := "http"
		domoClient := resty.New()
		config := &domopool_proto.Config{}

		domoClient.HostURL = scheme + "://192.168.11.183"
		domoClient.SetHeader("Content-Type", "application/json")
		domoClient.SetRetryCount(3)
		domoClient.SetRetryWaitTime(5 * time.Second)
		resp, err := domoClient.R().Get("/api/v1/config")
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal(resp.Body(), config)
		// fmt.Println(resp.String())
		// err = proto.Unmarshal(resp.Body(), config)
		if err != nil {
			fmt.Println(err)
		}

		switch getConfig {
		case "", "all":
			fmt.Println(config)
		case "mqtt":
			fmt.Println(config.Network.GetMqtt())
		}
		proto.Unmarshal(resp.Body(), config)

	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	configCmd.Flags().StringP("get", "g", "", "get config")
}
