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

var (
	cfgFile string
	web     embed.FS
)

const defaultIp = "192.168.11.183"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "domopool",
	Short: "The cli to manipulate a domopool box.",
	Run:   func(cmd *cobra.Command, args []string) {},
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
	rootCmd.PersistentFlags().String("box-host", defaultIp, "ip or hostname of the domopool box")
	rootCmd.PersistentFlags().Int("box-port", 80, "port of the domopool box")
	rootCmd.PersistentFlags().String("box-scheme", "http", "port of the domopool box")
	rootCmd.PersistentFlags().String("api-version", "v1", "api version")

	if err := viper.BindPFlag("boxHost", rootCmd.PersistentFlags().Lookup("box-host")); err != nil {
		log.Fatal().Err(err).Msg("Unable to set viper value")
	}
	if err := viper.BindPFlag("boxPort", rootCmd.PersistentFlags().Lookup("box-port")); err != nil {
		log.Fatal().Err(err).Msg("Unable to set viper value")
	}
	if err := viper.BindPFlag("boxScheme", rootCmd.PersistentFlags().Lookup("box-scheme")); err != nil {
		log.Fatal().Err(err).Msg("Unable to set viper value")
	}
	if err := viper.BindPFlag("api.version", rootCmd.PersistentFlags().Lookup("api-version")); err != nil {
		log.Fatal().Err(err).Msg("Unable to set viper value")
	}
	if err := viper.BindPFlag("api.path.config", rootCmd.PersistentFlags().Lookup("config")); err != nil {
		log.Fatal().Err(err).Msg("Unable to set viper value")
	}
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

	viper.SetDefault("boxHost", defaultIp)
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
