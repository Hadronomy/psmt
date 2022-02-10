package template

import (
	"fmt"

	initCmd "github.com/hadronomy/psmt/cli/cmd/template/init"
	"github.com/spf13/cobra"
)

func NewCmdTemplate() *cobra.Command {
	template := &cobra.Command{
		Use:   "template",
		Short: "Manages the psmt template system",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("WIP")
		},
	}
	template.AddCommand(initCmd.NewCmdInit())
	return template
}
