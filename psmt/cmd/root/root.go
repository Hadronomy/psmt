package root

import (
	"fmt"
	"os"

	templateCmd "github.com/hadronomy/psmt/psmt/cmd/template"
	"github.com/hadronomy/psmt/internal/build"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

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
	root.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.psmt.yaml)")

	root.AddCommand(templateCmd.NewCmdTemplate())
	return root
}

func Execute() {
	if err := NewCmdRoot(build.Version, build.Date).Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".psmt" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".psmt")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
