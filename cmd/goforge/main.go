package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/coderianx/goforge/internal/cli"
	"github.com/coderianx/goforge/internal/scaffold"
	"github.com/coderianx/goforge/internal/templates"
)

func main() {
	// arg kontrolü
	if len(os.Args) < 3 {
		fmt.Println("Usage: goforge new <project-name>")
		return
	}

	command := os.Args[1]
	projectName := os.Args[2]

	if command != "new" {
		fmt.Println("Unknown command:", command)
		return
	}

	// framework seç
	framework, err := cli.SelectFramework()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// embed içindeki template dizini (gin / fiber / echo / chi)
	templateDir := strings.ToLower(framework)

	// template'i embed.FS içinden kopyala
	err = scaffold.CopyDir(
		templates.FS,
		templateDir,
		projectName,
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	moduleName := projectName

	err = scaffold.RenderTemplate(
		templates.FS,
		templateDir+"/go.mod.tpl",
		projectName+"/go.mod",
		map[string]string{
			"ModuleName": moduleName,
		},
	)
	if err != nil {
		fmt.Println("Error generating go.mod:", err)
		return
	}

	fmt.Println("✅ Project created:", projectName)
	fmt.Println("👉 Framework:", framework)
	fmt.Println("👉 Next steps:")
	fmt.Println("   cd", projectName)
	fmt.Println("   go mod init", projectName)
	fmt.Println("   go run main.go")
}
