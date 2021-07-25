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

	"github.com/golgoth31/domopool-cli/internal/domoClient"
	"github.com/golgoth31/domopool-cli/internal/domoConfig"
	logger "github.com/golgoth31/domopool-cli/internal/log"
	"github.com/spf13/cobra"
)

// alarmsCmd represents the alarms command
var alarmsCmd = &cobra.Command{
	Use:   "alarms",
	Short: "View or reset alarms state.",
	ValidArgs: []string{
		"reset",
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 && args[0] == "reset" {
			client := domoClient.NewClient()
			logger.StdLog.Info().Msg("Reseting alarms")
			resp := client.Post("/api/v1/alarms/reset", nil)

			if resp.StatusCode() == 200 {
				time.Sleep(2 * time.Second)
			}
		}

		config := domoConfig.GetConfig()

		logger.StdLog.Info().Msgf("%v", config.GetAlarms())
	},
}

func init() {
	rootCmd.AddCommand(alarmsCmd)
}
