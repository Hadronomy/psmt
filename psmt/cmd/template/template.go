package template

import (
	"fmt"

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
	template.AddCommand(NewCmdInit())
	return template
}
