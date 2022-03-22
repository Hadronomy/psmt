package root

import (
	"fmt"
	"os"
	"path"

	"github.com/hadronomy/psmt/internal/build"
	daemonCmd "github.com/hadronomy/psmt/psmt/cmd/daemon"
	templateCmd "github.com/hadronomy/psmt/psmt/cmd/template"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Root *cobra.Command
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
	root.PersistentFlags().StringVar(&cfgFile, "config", "", "Overrides the default config file with the specified one")

	root.AddCommand(templateCmd.NewCmdTemplate())
	root.AddCommand(daemonCmd.NewCmdDaemon())
	return root
}

/*
Initializes and executes the cli program
*/
func Execute() {
	Root = NewCmdRoot(build.Version, build.Date)
	if err := Root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(path.Join(home, ".config", "psmt"))
		viper.SetConfigName("psmt")
		viper.AutomaticEnv()
	}
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
