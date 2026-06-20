package scaffold

import (
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"
)

func TestCopyDir(t *testing.T) {
	fsys := fstest.MapFS{
		"tpl/hello.gotmpl":   {Data: []byte(`package main`), Mode: 0644},
		"tpl/ping.go":         {Data: []byte(`package main`), Mode: 0644},
		"tpl/sub/data.txt":    {Data: []byte(`data`), Mode: 0644},
		"tpl/go.mod.tpl":      {Data: []byte(`module {{.ModuleName}}`), Mode: 0644},
	}

	dest, err := os.MkdirTemp("", "goforge-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dest)

	err = CopyDir(fsys, "tpl", dest)
	if err != nil {
		t.Fatal(err)
	}

	checks := []struct {
		path     string
		shouldExist bool
	}{
		{"hello.go", true},
		{"ping.go", true},
		{"sub/data.txt", true},
		{"hello.gotmpl", false},
		{"go.mod.tpl", false},
		{"go.mod", false},
	}

	for _, c := range checks {
		fullPath := filepath.Join(dest, c.path)
		_, err := os.Stat(fullPath)
		exists := err == nil
		if exists != c.shouldExist {
			t.Errorf("expected %s to exist=%v, got %v", c.path, c.shouldExist, exists)
		}
	}

	helloContent, err := os.ReadFile(filepath.Join(dest, "hello.go"))
	if err != nil {
		t.Fatal(err)
	}
	if string(helloContent) != "package main" {
		t.Errorf("expected 'package main', got '%s'", string(helloContent))
	}
}

func TestRenderTemplate(t *testing.T) {
	fsys := fstest.MapFS{
		"template.gotmpl": {Data: []byte(`module {{.ModuleName}}`), Mode: 0644},
	}

	dest, err := os.MkdirTemp("", "goforge-render-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dest)

	destPath := filepath.Join(dest, "go.mod")
	err = RenderTemplate(fsys, "template.gotmpl", destPath, map[string]string{
		"ModuleName": "testapp",
	})
	if err != nil {
		t.Fatal(err)
	}

	content, err := os.ReadFile(destPath)
	if err != nil {
		t.Fatal(err)
	}

	expected := "module testapp"
	if string(content) != expected {
		t.Errorf("expected '%s', got '%s'", expected, string(content))
	}
}

func TestAddDatabaseSupportPostgres(t *testing.T) {
	dir, err := os.MkdirTemp("", "goforge-db-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	err = AddDatabaseSupport(dir, "postgres")
	if err != nil {
		t.Fatal(err)
	}

	content, err := os.ReadFile(filepath.Join(dir, "database", "postgres.go"))
	if err != nil {
		t.Fatal(err)
	}

	if len(content) == 0 {
		t.Fatal("expected non-empty file")
	}
}

func TestAddDatabaseSupportSQLite(t *testing.T) {
	dir, err := os.MkdirTemp("", "goforge-db-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	err = AddDatabaseSupport(dir, "sqlite")
	if err != nil {
		t.Fatal(err)
	}

	content, err := os.ReadFile(filepath.Join(dir, "database", "sqlite.go"))
	if err != nil {
		t.Fatal(err)
	}

	if len(content) == 0 {
		t.Fatal("expected non-empty file")
	}
}

func TestAddDatabaseSupportInvalid(t *testing.T) {
	dir, err := os.MkdirTemp("", "goforge-db-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	err = AddDatabaseSupport(dir, "mysql")
	if err == nil {
		t.Fatal("expected error for unsupported database")
	}
}
