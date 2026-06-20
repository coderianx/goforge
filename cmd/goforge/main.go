package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/coderianx/goforge/internal/cli"
	"github.com/coderianx/goforge/internal/scaffold"
	"github.com/coderianx/goforge/internal/templates"
	"github.com/spf13/cobra"
)

var version = "0.2.0"

func main() {
	root := &cobra.Command{
		Use:     "goforge",
		Short:   "GoForge - Bootstrap Go projects with your preferred framework",
		Version: version,
	}

	root.AddCommand(&cobra.Command{
		Use:   "new <project-name>",
		Short: "Create a new Go project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			projectName := args[0]
			dbType, _ := cmd.Flags().GetString("db")

			framework, err := cli.SelectFramework()
			if err != nil {
				return fmt.Errorf("framework selection failed: %w", err)
			}

			templateDir := framework.DirName

			if err := scaffold.CopyDir(templates.FS, templateDir, projectName); err != nil {
				return fmt.Errorf("failed to copy template: %w", err)
			}

			if err := scaffold.RenderTemplate(
				templates.FS,
				templateDir+"/go.mod.tpl",
				projectName+"/go.mod",
				map[string]string{
					"ModuleName": projectName,
				},
			); err != nil {
				return fmt.Errorf("failed to generate go.mod: %w", err)
			}

			if dbType != "" {
				if err := scaffold.AddDatabaseSupport(projectName, dbType); err != nil {
					return fmt.Errorf("failed to add database support: %w", err)
				}
			}

			port := framework.Port

			fmt.Println("✅ Project created:", projectName)
			fmt.Println("👉 Framework:", framework.Name)
			fmt.Println("👉 Port:", port)
			fmt.Println("👉 Next steps:")
			fmt.Println("   cd", projectName)
			fmt.Println("   go mod tidy")
			fmt.Println("   go run .")

			if dbType != "" {
				fmt.Println("   Database:", strings.ToUpper(dbType[:1])+dbType[1:])
			}

			return nil
		},
	})

	newCmd := root.Commands()[0]
	newCmd.Flags().String("db", "", "Database driver (postgres, sqlite)")

	root.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List supported frameworks",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cli.PrintFrameworks()
			return nil
		},
	})

	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
