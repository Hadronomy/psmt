package init

import (
	"fmt"
	"os"
	"path"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

func NewCmdInit() *cobra.Command {
	init := &cobra.Command{
		Use:     "init",
		Short:   "Initializes a new template project",
		Example: "init",
		Run: func(cmd *cobra.Command, args []string) {
			pwd, _ := os.Getwd()
			var questions = []*survey.Question{
				{
					Name: "configFilename",
					Prompt: &survey.Select{
						Message: "Choose a config file type: ",
						Options: []string{
							"template.psmt.yaml",
							"template.yaml",
							"template.psmt.json",
							"template.json"},
					},
				},
			}
			answers := struct {
				ConfigFilename string `survey:"configFilename"`
			}{}
			err := survey.Ask(questions, &answers)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			os.Create(path.Join(pwd, answers.ConfigFilename))
		},
	}
	return init
}
