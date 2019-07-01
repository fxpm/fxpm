// Package servers provides a simplified management API around
// managing FX Server instances locally using the FXPM CLI.package
// servers
package servers

import (
	"fmt"

	"github.com/docker/docker/client"
	"github.com/fxpm/fxpm/contexts"
)

// ServerConfig provides a standard configuration to be used for
// management-related functionality with the CLI.
type ServerConfig struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Port        string `json:"port"`
	PortMode    string `json:"portMode"`
	RestartMode string `json:"restartMode"`
}

// Dock is the authenticated Docker client for use within the
// CLI to manipulate the connected Docker environment.
var Dock client.Client

// ContextSetup initializes the configuration of the current
// context. This is used to initialize important things in a
// specific order when running the CLI.
func ContextSetup(ctx contexts.Context) {
	fmt.Println("context loaded", ctx.Name)
}
