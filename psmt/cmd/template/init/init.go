package init

import (
	"fmt"
	"os"
	"path"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewCmdInit() *cobra.Command {
	init := &cobra.Command{
		Use:     "init",
		Short:   "Initializes a new template project",
		Example: "init",
		Run: func(cmd *cobra.Command, args []string) {
			pwd, _ := os.Getwd()
			var configFileNames = []string{
				"template.psmt.yaml",
				"template.yaml",
				"template.psmt.json",
				"template.json",
			}
			var questions = []*survey.Question{
				{
					Name: "config_filename",
					Prompt: &survey.Select{
						Message: "Choose a config file type: ",
						Options: configFileNames,
					},
				},
				{
					Name: "template_project_name",
					Prompt: &survey.Input{
						Message: "Template project name: ",
					},
				},
				{
					Name: "author",
					Prompt: &survey.Input{
						Message: "Author: ",
					},
				},
				{
					Name: "repo_url",
					Prompt: &survey.Input{
						Message: "Repository url: ",
					},
				},
			}
			answers := struct {
				ConfigFilename      string `survey:"config_filename"`
				TemplateProjectName string `survey:"template_project_name"`
				Author              string `survey:"author"`
				RepoUrl             string `survey:"repo_url"`
			}{}
			err := survey.Ask(questions, &answers)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			for _, configName := range configFileNames {
				if _, err := os.Stat(path.Join(pwd, configName)); err == nil {
					fmt.Println("This project already contains a psmt template configuration file")
					os.Exit(1)
				}
			}
			configPath := path.Join(pwd, answers.ConfigFilename)
			os.Create(configPath)
			templateConfig := viper.New()
			templateConfig.SetConfigFile(configPath)
			templateConfig.AutomaticEnv()
			templateConfig.SetEnvPrefix("psmt")
			templateConfig.Set("projectName", answers.TemplateProjectName)
			templateConfig.Set("author", answers.Author)
			templateConfig.Set("repo.url", answers.RepoUrl)
			templateConfig.Set("options", map[string]interface{}{
				"default": "hey",
			})
			templateConfig.WriteConfig()
		},
	}
	return init
}
