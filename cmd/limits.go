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

// limitsCmd represents the mqtt command
var limitsCmd = &cobra.Command{
	Use:   "limits",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Args: cobra.ExactValidArgs(1),
	ValidArgs: []string{
		"set",
	},
	Run: func(cmd *cobra.Command, args []string) {
		scheme := "http"
		domoClient := resty.New()
		limits := &domopool_proto.Config{}

		domoClient.HostURL = scheme + "://192.168.11.183"
		domoClient.SetRetryCount(3)
		domoClient.SetRetryWaitTime(5 * time.Second)

		resp := &resty.Response{}
		var err error

		if len(args) != 0 {
			fmt.Println("Setting limits")
			switch args[0] {
			case "set":
				limitsbody := &domopool_proto.Limits{
					WpMin:           0.2,
					WpMax:           2,
					PhMin:           0.0,
					PhMax:           0.0,
					ChMin:           0.0,
					ChMax:           0.0,
					WaitBeforeCh:    72,
					ChTempThreshold: 15,
					ChTempWaitReset: 14,
					Wp_0Derive:      0.03,
					TwMin:           1,
					TwMax:           30,
					TambMin:         0,
				}
				body, _ := proto.Marshal(limitsbody)
				resp, err = domoClient.
					R().
					SetBody(body).
					Post("/api/v1/limits/set")
				if err != nil {
					fmt.Println(err)
				}
				time.Sleep(5 * time.Second)
			}

			if resp.StatusCode() == 200 {
				time.Sleep(2 * time.Second)
			}
		}

		readResp, err := domoClient.R().Get("/api/v1/config")
		if err != nil {
			fmt.Println(err)
		}
		err = proto.Unmarshal(readResp.Body(), limits)
		if err != nil {
			fmt.Println(err)
		}
		if readResp.StatusCode() == 200 {
			fmt.Println(limits.GetLimits())
		}
	},
}

func init() {
	rootCmd.AddCommand(limitsCmd)

	limitsCmd.Flags().String("server", "s", "Help message for toggle")
}
