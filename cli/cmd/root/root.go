package root

import (
	"fmt"
	"os"

	markdownCmd "github.com/hadronomy/gocli/cli/cmd/markdown"
	"github.com/hadronomy/gocli/internal/build"
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
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if value, err := cmd.Flags().GetBool("build-date"); err == nil && value {
				fmt.Println(build.Date)
				os.Exit(0)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	root.AddCommand(markdownCmd.NewMarkdownCmd())
	root.PersistentFlags().Bool("build-date", false, "Prints the date and time when this binary was built")
	return root
}
