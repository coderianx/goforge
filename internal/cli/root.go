package cli

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

type Framework struct {
	Name     string
	DirName  string
	Port     string
	Database string
}

var Frameworks = []Framework{
	{Name: "Gin", DirName: "gin", Port: "8080"},
	{Name: "Fiber", DirName: "fiber", Port: "3000"},
	{Name: "Chi", DirName: "chi", Port: "8080"},
	{Name: "Echo", DirName: "echo", Port: "8080"},
	{Name: "Gorilla/Mux", DirName: "gorillamux", Port: "8080"},
	{Name: "Standard Library", DirName: "stdlib", Port: "8080"},
}

func SelectFramework() (Framework, error) {
	var choice Framework

	opts := make([]huh.Option[Framework], len(Frameworks))
	for i, f := range Frameworks {
		opts[i] = huh.NewOption(f.Name, f)
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[Framework]().
				Title("Select a Go framework:").
				Description("Choose your web framework").
				Options(opts...).
				Value(&choice),
		),
	)

	err := form.Run()
	return choice, err
}

func PrintFrameworks() {
	fmt.Println("Supported frameworks:")
	for _, f := range Frameworks {
		fmt.Printf("  - %s\n", f.Name)
	}
}
