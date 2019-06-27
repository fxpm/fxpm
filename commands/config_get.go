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

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// getCmd represents the get sub-command of config
var getCmd = &cobra.Command{
	Use:   "get <key>",
	Short: "Retrieves a configuration value for a specified key.",
	Long: `Retrieves a configuration value for a specified key.
For example:

config get fxpm.production
config get fxpm.debug
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var key string = args[0]
		if json, err := json.MarshalIndent(viper.Get(key), "", " "); err == nil {
			cmd.Printf("%s \n", json)

			return
		}

		cmd.Printf("%s \r\n", viper.GetString(args[0]))
	},
}

func init() {
	configCmd.AddCommand(getCmd)
}
