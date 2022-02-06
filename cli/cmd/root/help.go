package root

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

func rootHelpFunc(c *cobra.Command, s []string) {
	r := newTermRenderer()
	out, err := r.Render(generateHelp(c))
	if err != nil || out == "" {
		os.Exit(1)
	}
	fmt.Print(out)
}

func newTermRenderer() *glamour.TermRenderer {
	renderer, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithEmoji(),
	)
	return renderer
}

func generateHelp(c *cobra.Command) string {
	var help string
	help += "# " + c.CommandPath() + "\n"
	if c.Short != "" {
		help += c.Long + "\n\n"
	}
	help += "## Usage" + "\n"
	if c.UseLine() != "" {
		help += "\t" + c.UseLine() + "\n"
	}
	if c.HasAvailableSubCommands() {
		help += "\t" + c.CommandPath() + " [command]" + "\n\n"
		help += "## Avaidable Commands" + "\n"
		for _, subcommand := range c.Commands() {
			help += "- **" + subcommand.Name() + "**\t" + subcommand.Short + "\n"
		}
		help += "\n"
	}
	if c.HasAvailableFlags() {
		flagUsages := c.Flags().FlagUsages()
		help += "## Flags" + "\n"
		help += enlist(dedent(flagUsages)) + "\n"
	}
	if c.HasAvailableSubCommands() {
		help += "\n"
		help += "Use \"" + c.CommandPath() + " [command] --help for more information about a command"
	}
	return help
}

func dedent(s string) string {
	lines := strings.Split(s, "\n")
	var buf bytes.Buffer
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		fmt.Fprintln(&buf, strings.TrimPrefix(l, strings.Repeat(" ", 4)))
	}
	return strings.TrimSuffix(buf.String(), "\n")
}

func enlist(s string) string {
	lines := strings.Split(s, "\n")
	var buf bytes.Buffer
	for _, line := range lines {
		fmt.Fprintln(&buf, "- "+line)
	}
	return buf.String()
}
