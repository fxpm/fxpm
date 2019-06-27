// Copyright © 2019 Scott Plunkett <plunkets@aeoss.io>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package commands

import (
	"encoding/json"
	"fmt"

	"github.com/fxpm/fxpm/util"

	"github.com/fxpm/fxpm/logs"

	"gopkg.in/yaml.v2"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var outputType string = "yaml"

// configDumpCmd represents the configDump command
var configDumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump the current status of the configuration file.",
	Long: `Dump the current status of the configuration file.

The running configuration will be pullsed from the current
running configuration. Specifying an alternate configuration
file will result in teh alternate configuration being dumped.`,
	PreRun:  logs.CommandStarting,
	PostRun: logs.CommandEnded,
	Run: func(cmd *cobra.Command, args []string) {
		var types = []string{"yaml", "json"}
		var dump interface{}
		viper.Unmarshal(&dump)

		if !util.SliceContainsString(types, outputType) {
			fmt.Println("Invalid output specified. Showing yaml.")
		}

		var output []byte
		if outputType == "json" {
			output, _ = json.MarshalIndent(dump, "", "  ")
		} else {
			output, _ = yaml.Marshal(dump)
		}

		fmt.Printf("%s \n", output)
	},
}

func init() {
	configCmd.AddCommand(configDumpCmd)

	// Register the --out, -o flags for defining output type
	configDumpCmd.Flags().StringVarP(&outputType, "output", "o", "yaml", "defines the output type: yaml or json")
}
