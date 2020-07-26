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
		states := &domopool_proto.States{}

		domoClient.HostURL = scheme + "://192.168.11.183"
		domoClient.SetHeader("Content-Type", "application/json")
		domoClient.SetRetryCount(3)
		domoClient.SetRetryWaitTime(5 * time.Second)
		resp, err := domoClient.R().Get("/api/v1/states")
		if err != nil {
			fmt.Println(err)
		}
		// err = json.Unmarshal(resp.Body(), config)
		// fmt.Println(resp.String())
		err = proto.Unmarshal(resp.Body(), states)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(states)
	},
}

func init() {
	rootCmd.AddCommand(statesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
