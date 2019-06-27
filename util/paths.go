package util

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mitchellh/go-homedir"
)

// GetRootPath returns the path to the user's FXPM root.
func GetRootPath(dirs ...string) string {
	root, _ := homedir.Expand("~/.fxpm")

	if len(dirs) > 0 {
		return filepath.Join(root, strings.Join(dirs, strconv.QuoteRune(filepath.Separator)))
	}

	return root
}

// GetLogsPath returns the path to the user's FXPM logs directory.
func GetLogsPath(dirs ...string) string {
	logs, _ := homedir.Expand("~/.fxpm/logs")

	if len(dirs) > 0 {
		return filepath.Join(logs, strings.Join(dirs, strconv.QuoteRune(filepath.Separator)))
	}

	return logs
}

// GetTemplatesPath returns the path to the user's FXPM templates directory.
func GetTemplatesPath(dirs ...string) string {
	templates, _ := homedir.Expand("~/.fxpm/templates")

	if len(dirs) > 0 {
		return filepath.Join(templates, strings.Join(dirs, strconv.QuoteRune(filepath.Separator)))
	}

	return templates
}
