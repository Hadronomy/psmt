package root

import (
	markdownCmd "github.com/hadronomy/gocli/cli/cmd/markdown"
	"github.com/spf13/cobra"
)

func NewCmdRoot(version, buildDate string) *cobra.Command {
	root := &cobra.Command{
		Use:     "gocli",
		Version: version,
		Short:   "Just random testing with bubble tea and cobra",
		Long: `
	Yep... Just a silly little sandbox			
		`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	root.AddCommand(markdownCmd.NewMarkdownCmd())
	return root
}
