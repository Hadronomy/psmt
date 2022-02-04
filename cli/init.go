package cli

import (
	rootCmd "github.com/hadronomy/gocli/cli/cmd/root"
	"github.com/hadronomy/gocli/internal/build"
)

func Init() {
	version := build.Version
	date := build.Date
	var root = rootCmd.NewCmdRoot(version, date)
	root.Execute()
}
