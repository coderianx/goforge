package scaffold

import (
	"bytes"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"
)

func RenderTemplate(
	fsys fs.FS,
	templatePath string,
	destPath string,
	data any,
) error {
	tplBytes, err := fs.ReadFile(fsys, templatePath)
	if err != nil {
		return err
	}

	tpl, err := template.New("tpl").Parse(string(tplBytes))
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return err
	}

	return os.WriteFile(destPath, buf.Bytes(), 0644)
}
