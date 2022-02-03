package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(markdownCmd)
}

var markdownCmd = &cobra.Command{
	Use:   "markdown",
	Short: "Request user information in a lovely way",
	Run: func(cmd *cobra.Command, args []string) {
		in := `
# My little markdown
## Subtitle
Just a little [glamour](https://github.com/charmbracelet/glamour) test
`
		out, err := glamour.Render(in, "dark")
		if err != nil {
			os.Exit(1)
		}
		fmt.Print(out)
	},
}
