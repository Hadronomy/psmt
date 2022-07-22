package daemon

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

func NewCmdStart() *cobra.Command {
	start := &cobra.Command{
		Use:   "start [daemon-name]",
		Short: "Starts the specified daemon",
		Run: func(cmd *cobra.Command, args []string) {
			style := lipgloss.NewStyle().
				Bold(true).
				Blink(true).
				Foreground(lipgloss.Color("1"))
			fmt.Println(style.Render("WIP"))
		},
	}
	return start
}
