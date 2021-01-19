/*
Copyright © 2020 David Sabatie <david.sabatie@notrenet.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
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

// mqttCmd represents the mqtt command
var mqttCmd = &cobra.Command{
	Use:   "mqtt",
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
		"get",
	},
	Run: func(cmd *cobra.Command, args []string) {
		setServer, _ := cmd.Flags().GetString("server")
		scheme := "http"
		domoClient := resty.New()
		mqtt := &domopool_proto.Config{}

		domoClient.HostURL = scheme + "://192.168.11.183"
		domoClient.SetRetryCount(3)
		domoClient.SetRetryWaitTime(5 * time.Second)

		switch args[0] {
		case "enable":
			_, err := domoClient.
				R().
				Post("/api/v1/mqtt/enable")
			if err != nil {
				fmt.Println(err)
			}
			time.Sleep(5 * time.Second)
		case "disable":
			_, err := domoClient.
				R().
				Post("/api/v1/mqtt/disable")
			if err != nil {
				fmt.Println(err)
			}
			time.Sleep(5 * time.Second)
		case "set":
			mqtt.Network.Mqtt.Server = setServer
			body, _ := proto.Marshal(mqtt.Network.GetMqtt())
			_, err := domoClient.
				R().
				SetBody(body).
				Post("/api/v1/mqtt/set")
			if err != nil {
				fmt.Println(err)
			}
			time.Sleep(5 * time.Second)
		}

		resp, err := domoClient.R().Get("/api/v1/config")
		if err != nil {
			fmt.Println(err)
		}
		err = proto.Unmarshal(resp.Body(), mqtt)
		if err != nil {
			fmt.Println(err)
		}
		if resp.StatusCode() == 200 {
			fmt.Println(mqtt.Network.GetMqtt())
		}
	},
}

func init() {
	rootCmd.AddCommand(mqttCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mqttCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	mqttCmd.Flags().String("server", "s", "Help message for toggle")
}
