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

// alarmsCmd represents the alarms command
var alarmsCmd = &cobra.Command{
	Use:   "alarms",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	ValidArgs: []string{
		"reset",
	},
	Run: func(cmd *cobra.Command, args []string) {
		scheme := "http"
		domoClient := resty.New()
		config := &domopool_proto.Config{}

		domoClient.HostURL = scheme + "://192.168.11.183"
		domoClient.SetRetryCount(3)
		domoClient.SetRetryWaitTime(5 * time.Second)

		if len(args) != 0 && args[0] == "reset" {
			fmt.Println("Reseting alarms")
			resp, err := domoClient.
				R().
				Post("/api/v1/alarms/reset")
			if err != nil {
				fmt.Println(err)
			}

			if resp.StatusCode() == 200 {
				time.Sleep(2 * time.Second)
			}
		}

		resp, err := domoClient.R().Get("/api/v1/config")
		if err != nil {
			fmt.Println(err)
		}

		err = proto.Unmarshal(resp.Body(), config)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(config.GetAlarms())
	},
}

func init() {
	rootCmd.AddCommand(alarmsCmd)
}
