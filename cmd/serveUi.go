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
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"text/template"

	logger "github.com/golgoth31/domopool-cli/internal/log"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveUiCmd represents the serveUi command
var serveUiCmd = &cobra.Command{
	Use:   "serveUi",
	Short: "Server mode to access web UI.",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		e := echo.New()
		templateString, err := web.ReadFile("web/build/index.html")
		if err != nil {
			logger.StdLog.Fatal().Err(err).Msg("Unable to template index.html")
		}

		box := &IndexTemplate{
			DomopoolBoxHost:   viper.GetString("boxHost"),
			DomopoolBoxPort:   viper.GetInt("boxPort"),
			DomopoolBoxScheme: viper.GetString("boxScheme"),
		}
		t := &Template{
			templates: template.Must(template.New("index").Parse(string(templateString))),
		}
		e.Renderer = t
		web2root, err := fs.Sub(web, "web/build")
		if err != nil {
			logger.StdLog.Fatal().Err(err).Msg("Unable to read sub path")
		}
		e.GET(
			"/*",
			echo.WrapHandler(

				http.StripPrefix(
					"/",
					http.FileServer(
						http.FS(web2root),
					),
				),
			),
		)
		e.GET(
			"/",
			func(c echo.Context) error {
				return c.Render(http.StatusOK, "index", box)
			},
		)

		e.Logger.Fatal(
			e.Start(
				fmt.Sprintf(":%d", port),
			),
		)
	},
}

type IndexTemplate struct {
	DomopoolBoxHost   string
	DomopoolBoxPort   int
	DomopoolBoxScheme string
}
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func init() {
	rootCmd.AddCommand(serveUiCmd)
	serveUiCmd.Flags().IntP("port", "p", 8080, "port to listen on")
}
