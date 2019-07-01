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

	"github.com/fxpm/fxpm/contexts"
	"github.com/fxpm/fxpm/logs"

	"github.com/spf13/cobra"
)

var contextDescription string
var contextDockerScheme string
var contextDockerHost string

// contextsAddCmd represents the contextsAdd command
var contextsAddCmd = &cobra.Command{
	Use:   "add <name>",
	Short: "Creates a new context",
	Long: `Creates a new context to be used by FXPM and to allow saving
configuration values specific to a single context`,
	PreRun:  logs.CommandStarting,
	PostRun: logs.CommandEnded,
	Aliases: []string{"a"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name = args[0]
		Config := contexts.Create(name, contextDescription)

		if Config == nil {
			return
		}

		if err := Config.ReadInConfig(); err != nil {
			fmt.Println("The context config could not be read, but the Context has been created.")

			return
		}

		Config.Set("data.description", contextDescription)
		Config.Set("data.config.docker.scheme", contextDockerScheme)
		Config.Set("data.config.docker.host", contextDockerHost)

		fmt.Printf("The `%s` Context was created successfully. To switch to it, use `fxpm context use %s`. \n", name, Config.GetString("data.slug"))
	},
}

func init() {
	contextsCmd.AddCommand(contextsAddCmd)

	// Allows the user to specify the description of the Context to create without
	// requiring the use of the survey.
	contextsAddCmd.Flags().StringVarP(&contextDescription, "description", "d", "", "The description of the Context to create.")

	// Allows the user to specify the Docker scheme of the Context to create without
	// requiring the use of the survey.
	contextsAddCmd.Flags().StringVarP(&contextDockerScheme, "scheme", "s", "unix", "The Docker scheme of the Context to create.")

	// Allows the user to specify the DockerHost of the Context to create without
	// requiring the use of the survey.
	contextsAddCmd.Flags().StringVarP(&contextDockerHost, "host", "u", "/var/run/docker.sock", "The Docker host of the Context to create.")
}
