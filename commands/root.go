// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"fmt"
	"os"

	"github.com/fxpm/fxpm/contexts"
	"github.com/fxpm/fxpm/logs"
	"github.com/fxpm/fxpm/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// True when an action is forced via the config or via
// the -f, --force flags.
var FlagForced bool

// True when verbosity is enabled via the config or via
// the -v, --verbose flags.
var FlagVerbose bool

// True when debug is enabled via the config or via the
// -d, --debug flags.
var FlagDebug bool

// True when a manual configuration path is specified
// via the -c, --config flags.
var FlagConfig string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fxpm",
	Short: "A simple resource manager for FiveM's FX Server",
	Long: `A simple resource manager for FiveM's FX Server
	
FXPM provides a simple interface for managing packages installed
in an FX Server resources directory. FXPM aims to be simple to use
for both users and developers.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Specifiy a callback for once Cobra has been initialized
	cobra.OnInitialize(initConfig)

	// Force the execution of a specific command, even if the CLI
	// suggests against it.
	rootCmd.PersistentFlags().BoolVarP(&FlagForced, "force", "f", false, "force execution despite FXPM's best wishes")

	// Force verbose logging to happen for more detailed information
	// about running processes.
	rootCmd.PersistentFlags().BoolVarP(&FlagVerbose, "verbose", "v", false, "force verbose output")

	// Force debug logging to happen for more information that developers
	// might care about.
	rootCmd.PersistentFlags().BoolVarP(&FlagDebug, "debug", "x", false, "force debugging output")

	// Specify a non-default configuration path, this is passed in to
	// viper in order for it to understand where to load the file.
	rootCmd.PersistentFlags().StringVarP(&FlagConfig, "config", "c", "", "config file (default is config.yaml in $HOME/.fxpm or current dir)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	viper.SetDefault("fxpm.verbose", false)
	viper.SetDefault("fxpm.production", true)
	viper.SetDefault("fxpm.debug", false)

	viper.SetDefault("fxpm.templates.skipLocal", false)
	viper.SetDefault("fxpm.templates.localPath", util.GetTemplatesPath())

	viper.SetDefault("context", "default")

	viper.Set("fxpm.verbose", FlagVerbose)
	viper.Set("fxpm.debug", FlagDebug)

	if FlagConfig != "" {
		// Use config file from the flag.
		viper.SetConfigFile(FlagConfig)
	} else {
		var configPath = util.GetRootPath()
		var _, dirExistsErr = os.Stat(configPath)
		if os.IsNotExist(dirExistsErr) {
			configPathErr := os.MkdirAll(configPath, 0755)
			logsPathErr := os.MkdirAll(util.GetLogsPath(), 0755)
			if util.IsError(configPathErr) || util.IsError(logsPathErr) {
				fmt.Println(configPathErr, logsPathErr)

				return
			}
		}

		// Search config in home directory with name "config" (without extension).
		viper.AddConfigPath(configPath)
		viper.AddConfigPath(".")
		viper.SetConfigName("config")

		var configFilePath = util.GetRootPath("/config.yaml")
		var _, fileExistsErr = os.Stat(configFilePath)
		if os.IsNotExist(fileExistsErr) {
			var file, createErr = os.Create(configFilePath)
			defer file.Close()

			if createErr != nil {
				fmt.Println("Could not create config.yaml in $HOME/.fxpm")

				return
			}

			viper.WriteConfig()
		}
	}

	logs.SetupLog()

	viper.AutomaticEnv() // read in environment variables that match
	viper.SafeWriteConfig()

	contexts.Setup()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Failed to find config file. :(")
	}
}
