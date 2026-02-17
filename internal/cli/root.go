package cli

import "github.com/charmbracelet/huh"

func SelectFramework() (string, error) {
	var choice string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select a Go framework:").
				Options(
					huh.NewOption("Gin", "Gin"),
					huh.NewOption("Fiber", "Fiber"),
					huh.NewOption("Chi", "Chi"),
				).
				Value(&choice),
		),
	)

	err := form.Run()
	return choice, err
}
