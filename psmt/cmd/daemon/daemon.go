package daemon

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	initCmd "github.com/hadronomy/psmt/psmt/cmd/template/init"
	"github.com/spf13/cobra"
)

func NewCmdDaemon() *cobra.Command {
	daemon := &cobra.Command{
		Use:   "daemon",
		Short: "Manages the psmt daemon system",
		Run: func(cmd *cobra.Command, args []string) {
			style := lipgloss.NewStyle().
				Bold(true).
				Blink(true).
				Foreground(lipgloss.Color("1"))
			fmt.Println(style.Render("WIP"))
		},
	}
	daemon.AddCommand(initCmd.NewCmdInit())
	return daemon
}
