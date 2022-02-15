package init

import (
	"os"
	"path"

	"github.com/spf13/cobra"
)

func NewCmdInit() *cobra.Command {
	init := &cobra.Command{
		Use:     "init",
		Short:   "Initializes a new template project",
		Example: "init",
		Run: func(cmd *cobra.Command, args []string) {
			pwd, _ := os.Getwd()
			os.Create(path.Join(pwd, "template.psmt.yaml"))
		},
	}
	return init
}
