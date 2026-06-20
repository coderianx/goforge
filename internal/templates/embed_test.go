package templates

import (
	"io/fs"
	"strings"
	"testing"
)

func TestEmbeddedFrameworks(t *testing.T) {
	expected := []struct {
		dir      string
		files    []string
	}{
		{"gin", []string{"main.gotmpl", "hello.gotmpl", "ping.gotmpl", "go.mod.tpl", "go.sum", "Dockerfile", "docker-compose.yml", ".env.example"}},
		{"fiber", []string{"main.gotmpl", "hello.gotmpl", "ping.gotmpl", "go.mod.tpl", "go.sum", "Dockerfile", "docker-compose.yml", ".env.example"}},
		{"chi", []string{"main.gotmpl", "routes.gotmpl", "go.mod.tpl", "go.sum", "Dockerfile", "docker-compose.yml", ".env.example"}},
		{"echo", []string{"main.gotmpl", "hello.gotmpl", "ping.gotmpl", "go.mod.tpl", "go.sum", "Dockerfile", "docker-compose.yml", ".env.example"}},
		{"gorillamux", []string{"main.gotmpl", "hello.gotmpl", "ping.gotmpl", "go.mod.tpl", "go.sum", "Dockerfile", "docker-compose.yml", ".env.example"}},
		{"stdlib", []string{"main.gotmpl", "hello.gotmpl", "ping.gotmpl", "go.mod.tpl", "Dockerfile", "docker-compose.yml", ".env.example"}},
	}

	for _, fw := range expected {
		t.Run(fw.dir, func(t *testing.T) {
			for _, file := range fw.files {
				path := fw.dir + "/" + file
				info, err := fs.Stat(FS, path)
				if err != nil {
					t.Errorf("missing file %s: %v", path, err)
					continue
				}
				if info.Size() == 0 {
					t.Errorf("empty file: %s", path)
				}
			}
		})
	}
}

func TestEmbeddedNoExtraFiles(t *testing.T) {
	knownDirs := map[string]bool{
		"gin": true, "fiber": true, "chi": true,
		"echo": true, "gorillamux": true, "stdlib": true,
	}

	err := fs.WalkDir(FS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && path != "." {
			if !knownDirs[path] {
				t.Errorf("unexpected directory: %s", path)
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGoSumExistsForFrameworks(t *testing.T) {
	frameworksWithSum := []string{"gin", "fiber", "chi", "echo", "gorillamux"}
	for _, fw := range frameworksWithSum {
		info, err := fs.Stat(FS, fw+"/go.sum")
		if err != nil {
			t.Errorf("go.sum missing for %s: %v", fw, err)
			continue
		}
		if info.Size() == 0 {
			t.Errorf("go.sum empty for %s", fw)
		}
	}
}

func TestGoModTemplateSyntax(t *testing.T) {
	frameworks := []string{"gin", "fiber", "chi", "echo", "gorillamux", "stdlib"}
	for _, fw := range frameworks {
		data, err := fs.ReadFile(FS, fw+"/go.mod.tpl")
		if err != nil {
			t.Errorf("go.mod.tpl missing for %s: %v", fw, err)
			continue
		}
		content := string(data)
		if !strings.Contains(content, "{{.ModuleName}}") {
			t.Errorf("go.mod.tpl for %s missing ModuleName template variable", fw)
		}
	}
}
