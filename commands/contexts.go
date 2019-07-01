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
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// contextsCmd represents the contexts command
var contextsCmd = &cobra.Command{
	Use:   "contexts",
	Short: "Manage connectivity contexts within FXPM",
	Long: `The contexts command group offers the ability to setup
configuration items on a per-context basis, such as the Docker environment,
secrets, and keys.`,
	Aliases: []string{"ctx", "ct", "context"},
	Run: func(cmd *cobra.Command, args []string) {
		ctxMap := viper.GetStringMapString("contexts")

		var ctxArr [][]string
		for k, val := range ctxMap {
			var key = k

			if key == viper.GetString("context") {
				key = "* " + k
			} else {
				key = "  " + k
			}

			ctxArr = append(ctxArr, []string{key, val, "fxpm context use " + k})
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Key", "Path", "Command"})
		table.AppendBulk(ctxArr)
		table.SetCaption(true, "The currently active context is marked with an asterisk.")
		table.SetBorder(false)

		fmt.Println()
		table.Render()
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(contextsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// contextsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// contextsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
