package cli

import "github.com/AlecAivazis/survey/v2"

func SelectFramework() (string, error) {
	options := []string{
		"Gin",
		"Fiber",
	}

	var choice string
	prompt := &survey.Select{
		Message: "Select a Go framework:",
		Options: options,
	}

	err := survey.AskOne(prompt, &choice)
	return choice, err
}
