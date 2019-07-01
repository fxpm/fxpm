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

	"github.com/spf13/cobra"
)

var name string
var port string

// serversCreateCmd represents the serversCreate command
var serversCreateCmd = &cobra.Command{
	Use:   "create <dir>",
	Short: "Create an FX Server instance locally with FXPM",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Aliases: []string{"c", "add", "a"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serversCreate called")
	},
}

func init() {
	serversCmd.AddCommand(serversCreateCmd)

	// Specify the name of the FXPM server through flags, rather than through
	// the wizard
	serversCreateCmd.Flags().StringVarP(&name, "name", "n", "", "Specify a name for the server instance")

	// Specify a specific port to use for the FXPM server through flags, rather
	// than through the wizard
	serversCreateCmd.Flags().StringVarP(&port, "port", "p", "", "Sepcify a specific port to listen on")
}
