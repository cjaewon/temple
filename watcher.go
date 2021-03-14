package temple

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func watchGlob(pattern string) error {
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return err
	}

	watch(matches...)

	return nil
}

func watchFS(fsys fs.FS, patterns []string) error {
	var filenames []string

	for _, pattern := range patterns {
		list, err := fs.Glob(fsys, pattern)
		if err != nil {
			return err
		}
		if len(list) == 0 {
			return fmt.Errorf("template: pattern matches no files: %#q", pattern)
		}
		filenames = append(filenames, list...)
	}

	watch(filenames...)
	return nil
}

func watch(filenames ...string) {

}
