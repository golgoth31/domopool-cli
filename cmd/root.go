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
	"embed"
	"fmt"
	"os"

	logger "github.com/golgoth31/domopool-cli/internal/log"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var web embed.FS

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ardipool-cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(webFiles embed.FS) {
	web = webFiles
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.domopool-cli.yaml)")
	rootCmd.PersistentFlags().Bool("debug", false, "debug")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	logLevel, err := rootCmd.Flags().GetBool("debug")
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	if logLevel {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	logger.Initialize()

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			logger.StdLog.Fatal().Err(err).Msg("")
		}

		viper.SetConfigFile(fmt.Sprintf("%s/.domopool-cli.yaml", home))
	}
	viper.SetConfigType("yaml")

	viper.AutomaticEnv() // read in environment variables that match

	viper.SetDefault("boxIP", "192.168.11.183")
	viper.SetDefault("boxScheme", "http")
	viper.SetDefault("api.version", "v1")
	viper.SetDefault("api.path.config", "config")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		logger.StdLog.Debug().Msgf("Using config file: %s", viper.ConfigFileUsed())
	} else {
		logger.StdLog.Debug().Msgf("Saving config file")
		if err := viper.WriteConfig(); err != nil {
			logger.StdLog.Debug().Msgf("Error saving file")
		}
	}
}
