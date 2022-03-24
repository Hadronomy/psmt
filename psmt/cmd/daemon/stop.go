package daemon

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

func NewCmdStop() *cobra.Command {
	stop := &cobra.Command{
		Use:   "stop [daemon-name]",
		Short: "Stops the specified daemon",
		Run: func(cmd *cobra.Command, args []string) {
			style := lipgloss.NewStyle().
				Bold(true).
				Blink(true).
				Foreground(lipgloss.Color("1"))
			fmt.Println(style.Render("WIP"))
		},
	}
	return stop
}