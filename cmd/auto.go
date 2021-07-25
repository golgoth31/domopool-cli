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
	"time"

	"github.com/golgoth31/domopool-cli/internal/domoClient"
	"github.com/golgoth31/domopool-cli/internal/domoConfig"
	logger "github.com/golgoth31/domopool-cli/internal/log"
	"github.com/spf13/cobra"
)

// autoCmd represents the filter command
var autoCmd = &cobra.Command{
	Use:   "auto",
	Short: "Set automatic or recover mode.",
	Run: func(cmd *cobra.Command, args []string) {
		client := domoClient.NewClient()
		logger.StdLog.Info().Msg("Setting domopool in automatic mode")

		resp := client.Post("/api/v1/auto", nil)

		recover, _ := cmd.Flags().GetBool("recover")
		if recover {
			resp = client.Post("/api/v1/recover", nil)
			logger.StdLog.Info().Msg("Forcing recover mode")
		}

		if resp.StatusCode() == 200 {
			time.Sleep(2 * time.Second)
			config := domoConfig.GetConfig()

			logger.StdLog.Info().Msgf("%v", config.GetPump())
			logger.SuccessLog.Info().Msg("Done")
		}
	},
}

func init() {
	rootCmd.AddCommand(autoCmd)
	autoCmd.Flags().BoolP("recover", "r", false, "Put pool in recover mode (reopening before summer)")
}
