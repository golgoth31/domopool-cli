/*
Copyright © 2021 David Sabatie <david.sabatie@notrenet.com>

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
	"github.com/golgoth31/domopool-cli/configs"
	"github.com/golgoth31/domopool-cli/internal/domoClient"
	logger "github.com/golgoth31/domopool-cli/internal/log"
	domopool_proto "github.com/golgoth31/domopool-proto"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		logger.StdLog.Info().Msgf("Version: %v, build from: %v, on: %v\n", configs.Version, configs.GitCommit, configs.BuildDate)

		infos := &domopool_proto.Infos{}
		client := domoClient.NewClient()
		resp := client.Get("/")

		if err := proto.Unmarshal(resp.Body(), infos); err != nil {
			logger.StdLog.Fatal().Err(err).Msg("Unable to unmarchal proto")
		}

		logger.StdLog.Info().Msgf("Board compile time => %s", infos.GetCompile())
		logger.StdLog.Info().Msgf("Board type => %s", infos.GetBoardName())
		logger.StdLog.Info().Msgf("Board component versions => %v", infos.GetVersions())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
