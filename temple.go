package temple

import (
	"path/filepath"
	"text/template"
)

func NewHot(pattern string, t *template.Template) (*Watcher, error) {
	watcher := newWatcher(t)

	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	for _, match := range matches {
		watcher.Add(match)
	}

	return watcher, nil
}
