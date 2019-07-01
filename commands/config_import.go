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
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/fxpm/fxpm/logs"

	"github.com/fxpm/fxpm/util"

	"github.com/mitchellh/go-homedir"

	"github.com/spf13/cobra"
)

// configImportCmd represents the configImport command
var configImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import FXPM configuration from a file",
	Long: `Import FXPM configuration from a .yaml, .yml,
or .json file.
	
Provide the relative path to a file in either yaml or
json format to import the file and overwrite the existing
configuration. Specifying an alternative configuration will
overwrite the alternative configuration instead of the global
configuration.`,
	Args:    cobra.ExactArgs(1),
	Aliases: []string{"i"},
	Run: func(cmd *cobra.Command, args []string) {
		var possibleTypes []string = []string{".json", ".yaml", ".yml"}
		path, _ := homedir.Expand(args[0])
		fileType := filepath.Ext(path)

		if fileType == "" {
			cmd.PrintErrln("The file type could not be inferred. Please use a file of an accepted type with a valid extension.")
			cmd.PrintErrf("Valid types include: %s \n", strings.Join(possibleTypes, ", "))

			return
		}

		if !util.SliceContainsString(possibleTypes, fileType) {
			cmd.PrintErrln("The path specified does not point to an acceptable file type. Please use a valid file type.")
			cmd.PrintErrf("Valid types include: %s \n", strings.Join(possibleTypes, ", "))

			return
		}

		// Fetch file contents
		fetchContents := logs.NewCommandProcess("fetch contents", "Fetching file contents...")

		_, e := ioutil.ReadFile(path)
		logs.ErrorIf(e, "FXPM couldn't read the file contents.")

		fetchContents.Done("Contents have been read from the import file.")

		// Save the contents of the import file to the config path.
		saveContents := logs.NewCommandProcess("save contents", "Converting contents to FXPM configuration format")
		saveContents.Done("Contents have been saved.")
	},
}

func init() {
	configCmd.AddCommand(configImportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configImportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configImportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
