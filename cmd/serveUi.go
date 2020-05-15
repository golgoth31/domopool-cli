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
	"log"
	"net"
	"net/http"
	"text/template"

	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

// serveUiCmd represents the serveUi command
var serveUiCmd = &cobra.Command{
	Use:   "serveUi",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		e := echo.New()
		indexBox, err := rice.FindBox("../../build")
		if err != nil {
			log.Fatal(err)
		}
		jsBox, err := rice.FindBox("../../build/bundle")
		if err != nil {
			log.Fatal(err)
		}

		templateString, err := indexBox.String("index.html")
		if err != nil {
			log.Fatal(err)
		}
		// parse and execute the template
		// tmplIndex, err := template.New("index").Parse(templateString)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		box_ip, err := cmd.Flags().GetIP("box-ip")
		if err != nil {
			log.Fatal(err)
		}
		box_port, err := cmd.Flags().GetInt("box-port")
		if err != nil {
			log.Fatal(err)
		}
		box := &IndexTemplate{
			DomopoolIP:   box_ip.String(),
			DomopoolPort: box_port,
		}
		t := &Template{
			templates: template.Must(template.New("index").Parse(templateString)),
		}
		e.Renderer = t
		// bundle, err := jsBox.String("bundle.js")
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println(bundle)
		e.GET(
			"/bundle/*",
			echo.WrapHandler(
				http.StripPrefix("/bundle/",
					http.FileServer(
						jsBox.HTTPBox(),
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
	DomopoolIP   string
	DomopoolPort int
}
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func init() {
	rootCmd.AddCommand(serveUiCmd)

	defaultIp := net.ParseIP("127.0.0.1")
	serveUiCmd.Flags().IntP("port", "p", 8080, "port to listen on")
	serveUiCmd.Flags().IPP("box-ip", "i", defaultIp, "ip of the domopool box")
	serveUiCmd.Flags().Int("box-port", 80, "port of the domopool box")
}
