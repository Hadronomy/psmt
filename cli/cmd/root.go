package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gocli",
	Short: "Just random testing with bubbles and cobra",
	Long: `
Yep... Just a silly little sandbox			
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// Empty for now
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
