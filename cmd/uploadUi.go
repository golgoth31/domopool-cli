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
	"fmt"
	"log"
	"os"
	"time"

	rice "github.com/GeertJohan/go.rice"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/cobra"
)

type uploadFile struct {
	filename string
	data     []byte
}

// uploadUiCmd represents the uploadUi command
var uploadUiCmd = &cobra.Command{
	Use:   "uploadUi",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		scheme := "http"
		domoClient := resty.New()

		domoClient.HostURL = scheme + "://192.168.11.183"
		domoClient.SetRetryCount(3)
		domoClient.SetRetryWaitTime(5 * time.Second)

		uiBox, err := rice.FindBox("../web/build")
		if err != nil {
			log.Fatal(err)
		}
		index, err := uiBox.Bytes("index.html")
		if err != nil {
			fmt.Println("can't read index file")
			os.Exit(1)
		}
		bundle, err := uiBox.Bytes("bundle.js")
		if err != nil {
			fmt.Println("can't read index file")
			os.Exit(1)
		}

		fmt.Println("uploading index")
		resp, err := domoClient.
			R().
			SetFileReader("", "index.html", bytes.NewReader(index)).
			SetFormData(map[string]string{
				"filename": "index.html",
			}).
			SetContentLength(true).
			Post("/ui/upload")
		if err != nil {
			fmt.Println(err)
		}

		if resp.StatusCode() == 200 {
			fmt.Println("uploading bundle")
			resp, err := domoClient.
				R().
				SetFileReader("", "bundle.js", bytes.NewReader(bundle)).
				SetFormData(map[string]string{
					"filename": "bundle.js",
				}).
				SetContentLength(true).
				Post("/ui/upload")
			if err != nil {
				fmt.Println(err)
			}
			if resp.StatusCode() == 200 {
				fmt.Println("upload ok")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(uploadUiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadUiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadUiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
