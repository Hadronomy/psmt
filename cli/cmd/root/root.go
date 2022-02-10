package root

import (
	"fmt"
	"os"

	templateCmd "github.com/hadronomy/psmt/cli/cmd/template"
	"github.com/hadronomy/psmt/internal/build"
	"github.com/spf13/cobra"
)

func NewCmdRoot(version, buildDate string) *cobra.Command {
	root := &cobra.Command{
		Use:     "psmt",
		Version: version,
		Short:   "Please save my time :cry:",
		Long: `
A command line toolbox to save you, your time`,
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

	root.SetHelpFunc(rootHelpFunc)
	root.PersistentFlags().Bool("build-date", false, "Prints the date and time when this binary was built")

	root.AddCommand(templateCmd.NewCmdTemplate())
	return root
}
