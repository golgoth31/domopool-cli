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
	"github.com/golgoth31/domopool-cli/internal/domoConfig"
	logger "github.com/golgoth31/domopool-cli/internal/log"
	"github.com/spf13/cobra"
)

// metricsCmd represents the metrics command
var metricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Show metrics.",
	Run: func(cmd *cobra.Command, args []string) {
		config := domoConfig.GetConfig()
		logger.StdLog.Info().Msgf("%v", config.GetMetrics())
	},
}

func init() {
	rootCmd.AddCommand(metricsCmd)
}
