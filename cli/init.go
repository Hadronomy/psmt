package cli

import (
	rootCmd "github.com/hadronomy/psmt/cli/cmd/root"
	"github.com/hadronomy/psmt/internal/build"
)

func Init() {
	version := build.Version
	date := build.Date
	root := rootCmd.NewCmdRoot(version, date)
	root.Execute()
}
