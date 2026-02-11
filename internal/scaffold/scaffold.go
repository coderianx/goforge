package scaffold

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func CopyDir(fsys fs.FS, srcDir, destDir string) error {
	return fs.WalkDir(fsys, srcDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// .tpl dosyalarını kopyalama
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".tpl") {
			return nil
		}

		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(destDir, relPath)

		if d.IsDir() {
			return os.MkdirAll(targetPath, 0755)
		}

		srcFile, err := fsys.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		destFile, err := os.Create(targetPath)
		if err != nil {
			return err
		}
		defer destFile.Close()

		_, err = io.Copy(destFile, srcFile)
		return err
	})
}
