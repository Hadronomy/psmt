package init

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCmdInit() *cobra.Command {
	init := &cobra.Command{
		Use:     "init",
		Short:   "Initializes a new template project",
		Example: "init",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("WIP")
		},
	}

	return init
}
