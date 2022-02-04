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
# gocli
Yep... Just a silly little sandbox

## Usage:
	gocli [flags]
	gocli [command]

## Available Commands:
- **completion**  Generate the autocompletion script for the specified shell
- **help**        Help about any command
- **markdown**    Request user information in a lovely way
- **version**     Prints the tool version

## Flags:
- -h, --help      help for gocli
- -v, --version   version for gocli

Use "gocli [command] --help" for more information about a command.
`
		out, err := glamour.Render(in, "dark")
		if err != nil {
			os.Exit(1)
		}
		fmt.Print(out)
	},
}
