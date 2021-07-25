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
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/golgoth31/domopool-cli/internal/domoClient"
	"github.com/golgoth31/domopool-cli/internal/domoConfig"
	logger "github.com/golgoth31/domopool-cli/internal/log"
	domopool_proto "github.com/golgoth31/domopool-proto"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

// mqttCmd represents the mqtt command
var mqttCmd = &cobra.Command{
	Use:   "mqtt",
	Short: "Enable/disable or set mqtt.",
	Args:  cobra.ExactValidArgs(1),
	ValidArgs: []string{
		"enable",
		"disable",
		"set",
	},
	Run: func(cmd *cobra.Command, args []string) {
		setServer, _ := cmd.Flags().GetString("server")
		client := domoClient.NewClient()
		resp := &resty.Response{}

		switch args[0] {
		case "enable":
			resp = client.Post("/api/v1/mqtt/enable", nil)
			time.Sleep(5 * time.Second)
		case "disable":
			resp = client.Post("/api/v1/mqtt/disable", nil)
			time.Sleep(5 * time.Second)
		case "set":
			mqttbody := &domopool_proto.Mqtt{
				Server: setServer,
			}
			body, _ := proto.Marshal(mqttbody)
			resp = client.Post("/api/v1/mqtt/set", body)
			time.Sleep(5 * time.Second)
		}

		if resp.StatusCode() == 200 {
			time.Sleep(2 * time.Second)
			config := domoConfig.GetConfig()
			logger.StdLog.Info().Msgf("%v", config.Network.GetMqtt())
		}
	},
}

func init() {
	rootCmd.AddCommand(mqttCmd)

	mqttCmd.Flags().String("server", "s", "Help message for toggle")
}
