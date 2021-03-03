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
	domopool_proto "github.com/golgoth31/domopool-proto"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

// statesCmd represents the states command
var statesCmd = &cobra.Command{
	Use:   "states",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		scheme := "http"
		domoClient := resty.New()
		states := &domopool_proto.Config{}

		domoClient.HostURL = scheme + "://192.168.11.183"
		domoClient.SetHeader("Content-Type", "application/json")
		domoClient.SetRetryCount(3)
		domoClient.SetRetryWaitTime(5 * time.Second)
		resp, err := domoClient.R().Get("/api/v1/config")
		if err != nil {
			fmt.Println(err)
		}
		err = proto.Unmarshal(resp.Body(), states)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(states.GetStates())
	},
}

func init() {
	rootCmd.AddCommand(statesCmd)
}
