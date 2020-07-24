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
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/golgoth31/domopool-cli/internal"

	domopool_proto "github.com/golgoth31/domopool-proto"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
)

// testUiCmd represents the testServer command
var testUiCmd = &cobra.Command{
	Use:   "testUi",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		scheme := "http"
		domoClient := resty.New()
		config := &domopool_proto.Config{}

		domoClient.HostURL = scheme + "://192.168.11.183"
		domoClient.SetHeader("Content-Type", "application/json")
		domoClient.SetRetryCount(3)
		domoClient.SetRetryWaitTime(5 * time.Second)
		resp, err := domoClient.R().Get("/api/v1/config")
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal(resp.Body(), config)
		// fmt.Println(resp.String())
		// err = proto.Unmarshal(resp.Body(), config)
		if err != nil {
			fmt.Println(err)
		}

		e := echo.New()
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{
				http.MethodGet,
				http.MethodHead,
				http.MethodPut,
				http.MethodPatch,
				http.MethodPost,
				http.MethodDelete,
			},
		}))
		e.GET("/api/v1/config", func(c echo.Context) error {
			// resp := domopool_proto.Config{
			// 	Global: &domopool_proto.Global{
			// 		AckTone: 4000,
			// 	},
			// }
			return c.JSON(http.StatusOK, config)
		})
		e.POST("/config", func(c echo.Context) error {
			u := internal.Aconfig{}
			if err := c.Bind(u); err != nil {
				return err
			}
			log.Printf("%+v", u)
			e.Logger.Info(u)
			return c.JSON(http.StatusOK, u)
		})
		port, _ := cmd.Flags().GetInt("port")
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
	},
}

func init() {
	rootCmd.AddCommand(testUiCmd)
	testUiCmd.Flags().IntP("port", "p", 8080, "port to listen on")
}
