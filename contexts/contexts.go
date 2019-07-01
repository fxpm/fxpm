// Package contexts provides the means for switching configuration
// contexts in order to connect to different endpoints and backends,
// such as a remote Docker server.
package contexts

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/fxpm/fxpm/util"
	"github.com/gosimple/slug"
	"github.com/spf13/viper"
)

// Viper stores the Viper instance of the current contest to be used
// by the CLI.
var Viper *viper.Viper

// ContextDataConfig holds information about the configuration file to
// be read/written from for storing context-related items.
type ContextDataConfig struct {
	Docker ContextDataDocker `json:"docker"`
}

// ContextDataDocker holds the Docker-related configuration for generating
// the Docker SDK client.
type ContextDataDocker struct {
	Scheme    string            `json:"scheme"`
	Host      string            `json:"host"`
	Headers   map[string]string `json:"headers"`
	Timeout   time.Duration     `json:"timeout"`
	Version   string            `json:"version"`
	TLSClient struct {
		CACertPath string `json:"caCertPath"`
		CertPath   string `json:"certPath"`
		KeyPath    string `json:"keyPath"`
	}
}

// ContextData provides the means for the CLI to read information
// from the currently selected Context.
type ContextData struct {
	Slug        string                 `json:"slug"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Path        string                 `json:"path"`
	Meta        map[string]interface{} `json:"meta"`
	Config      ContextDataConfig      `json:"config"`
}

// Init reads the current configuration from the global Viper instance,
// and initalizes the selected context for the CLI run.
func Init() {

}

// Setup is called once at the beginning of the FXPM execution to ensure
// that necessary items are created for the CLI to execute properly.
func Setup() {
	var contextPath = util.GetRootPath("contexts")
	_, contextsMissing := os.Stat(contextPath)

	if util.IsError(contextsMissing) {
		err := os.MkdirAll(contextPath, 0755)

		if util.IsError(err) {
			fmt.Printf("Unable to setup Contexts. Please create the the following database structure and retry. \n")
			fmt.Println(contextPath)

			return
		}
	}

	var defaultContextPath = util.GetRootPath("contexts", "default.yaml")
	_, defaultMissing := os.Stat(defaultContextPath)

	if util.IsError(defaultMissing) {
		var defaultFile = `data:
	slug: default
	name: Default
	description: ""
	path: /Users/scott/.fxpm/contexts/default.yaml
	meta: {}
	config:
		docker:
		scheme: unix
		host: /var/run/docker.sock
		headers: {}
		timeout: 0s
		version: ""
		tlsclient:
			cacertpath: ""
			certpath: ""
			keypath: ""`
		var defaultContext = []byte{}
		copy(defaultContext, defaultFile)
		err := ioutil.WriteFile(defaultContextPath, defaultContext, 0755)

		if util.IsError(err) {
			fmt.Printf("Unable to setup Contexts. Please contaxt the FXPm developers or retry later. \n")
			fmt.Println(err.Error())

			return
		}

		viper.Set("contexts.default", defaultContextPath)
		viper.WriteConfig()
	}
}

// Create ...
func Create(name string, description string) *viper.Viper {
	var contextSlug = slug.Make(name)
	var contextPath = util.GetRootPath("contexts", contextSlug+".yaml")

	if viper.GetString("contexts."+contextSlug) != "" {
		fmt.Printf("The context `%s` already exists. Please choose a new name. \n", contextSlug)

		return nil
	}

	_, contextMissing := os.Stat(contextPath)
	if util.IsError(contextMissing) {
		_, err := os.Create(contextPath)

		if util.IsError(err) {
			fmt.Println(err.Error())
			fmt.Println("The context could not be created. The context file could not be created.")

			return nil
		}
	}

	Viper = viper.New()
	Viper.SetConfigFile(contextPath)

	contextData := ContextData{
		Slug:        contextSlug,
		Name:        name,
		Description: description,
		Path:        contextPath,
		Meta:        make(map[string]interface{}),
		Config: ContextDataConfig{
			Docker: ContextDataDocker{
				Scheme: "unix",
				Host:   "/var/run/docker.sock",
			},
		},
	}

	Viper.Set("data", contextData)
	Viper.WriteConfig()

	viper.Set("contexts."+contextSlug, contextPath)
	viper.WriteConfig()

	return Viper
}

// Config returns the Viper instance for the initialized Context which
// is then made avaialble to commands requiring the Context to execute.
func Config() *viper.Viper {
	currentContext := viper.GetString("context")
	currentContextPath := viper.GetString("contexts." + currentContext)

	if currentContext == "" || currentContextPath == "" {
		fmt.Println("Could not fetch the current Context. Check your config.yaml file.")

		return nil
	}

	Viper = viper.New()
	Viper.SetConfigFile(currentContextPath)

	if err := Viper.ReadInConfig(); err == nil {
		fmt.Printf("Could not read Context config. Check that .yaml file exists at `%s`. \n", currentContextPath)
	}

	return Viper
}
