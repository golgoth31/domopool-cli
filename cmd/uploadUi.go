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
	"bytes"

	"github.com/golgoth31/domopool-cli/internal/domoClient"
	logger "github.com/golgoth31/domopool-cli/internal/log"
	"github.com/spf13/cobra"
)

// uploadUiCmd represents the uploadUi command
var uploadUiCmd = &cobra.Command{
	Use:   "uploadUi",
	Short: "Upload updated UI to domopool box.",
	Run: func(cmd *cobra.Command, args []string) {
		client := domoClient.NewClient()

		index, err := web.ReadFile("web/build/index.html")
		if err != nil {
			logger.StdLog.Fatal().Err(err).Msg("Unable to read file")
		}
		bundle, err := web.ReadFile("web/build/bundle.js")
		if err != nil {
			logger.StdLog.Fatal().Err(err).Msg("Unable to read file")
		}
		favicon, err := web.ReadFile("web/build/favicon.ico")
		if err != nil {
			logger.StdLog.Fatal().Err(err).Msg("Unable to read file")
		}

		logger.StdLog.Info().Msg("uploading index")
		resp, err := client.Client.
			R().
			SetFileReader("", "index.html", bytes.NewReader(index)).
			SetFormData(map[string]string{
				"filename": "index.html",
			}).
			SetContentLength(true).
			Post("/ui/upload")
		if err != nil {
			logger.StdLog.Fatal().Err(err).Msg("Unable to request box")
		}

		if resp.StatusCode() == 200 {
			logger.StdLog.Info().Msg("uploading favicon")
			resp, err = client.Client.
				R().
				SetFileReader("", "favicon.ico", bytes.NewReader(favicon)).
				SetFormData(map[string]string{
					"filename": "favicon.ico",
				}).
				SetContentLength(true).
				Post("/ui/upload")
			if err != nil {
				logger.StdLog.Fatal().Err(err).Msg("Unable to request box")
			}
			if resp.StatusCode() == 200 {
				logger.StdLog.Info().Msg("uploading bundle")
				resp, err := client.Client.
					R().
					SetFileReader("", "bundle.js", bytes.NewReader(bundle)).
					SetFormData(map[string]string{
						"filename": "bundle.js",
					}).
					SetContentLength(true).
					Post("/ui/upload")
				if err != nil {
					logger.StdLog.Fatal().Err(err).Msg("Unable to request box")
				}
				if resp.StatusCode() == 200 {
					logger.StdLog.Info().Msg("upload ok")
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(uploadUiCmd)
}
