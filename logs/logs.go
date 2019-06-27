// Package logs provides the logging mechanisms of the FXPM
// CLI utility, including helper methods for command
// logging and step logging.
package logs

import (
	"fmt"
	"os"
	"time"

	"github.com/fxpm/fxpm/util"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log stores the zap.Logger instance of the global
// logger.
var Log *zap.Logger

// StartTime stored the starting time of a command's
// execution.
var StartTime time.Time

// BuildLogger is a function to build a zap.Logger
// instance using FXPM standard config.
func BuildLogger() (*zap.Logger, error) {
	var LogPath = util.GetLogsPath("standard.log")

	cfg := zap.NewProductionConfig()
	cfg.Development = !viper.GetBool("fxpm.production")
	cfg.OutputPaths = []string{LogPath, "stdout"}
	cfg.Encoding = "console"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.DisableStacktrace = viper.GetBool("fxpm.production")
	cfg.DisableCaller = viper.GetBool("fxpm.production")

	if viper.GetBool("fxpm.debug") {
		cfg.Level.SetLevel(zapcore.DebugLevel)
	} else {
		if viper.GetBool("fxpm.verbose") {
			cfg.Level.SetLevel(zapcore.InfoLevel)
		} else {
			cfg.Level.SetLevel(zapcore.WarnLevel)
		}
	}

	return cfg.Build()
}

// SetupLog is used to setup global logging and set the
// global Log variable.
func SetupLog() {
	localLog, e := BuildLogger()

	if util.IsNotError(e) {
		Log = localLog

		return
	}

	fmt.Println("Error starting logger. FXPM will exit now.")
	fmt.Printf("Logger encountered error, %s", e)
	os.Exit(1)
}

// ErrorIf provides a Simplified way of checking if an error was
// found and logging a message as Error if so.
func ErrorIf(e error, message string) {
	if Log == nil {
		fmt.Println("Couldn't write to Log, Log was not available.")
		os.Exit(1)

		return
	}

	if util.IsError(e) {
		fmt.Printf("Error: %s \r\n", e.Error())
		Log.Error(message, zap.String("error", e.Error()))

		return
	}
}
