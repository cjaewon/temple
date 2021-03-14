package temple

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	FSWatcher *fsnotify.Watcher
	template  *Template
}

func newWatcher(t *Template) *Watcher {
	w := Watcher{}

	fsWatcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}

	w.template = t
	w.FSWatcher = fsWatcher

	go w.WatchWorker()

	return &w
}

func (w *Watcher) WatchGlob(pattern string) error {
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return err
	}

	if err := w.Watch(matches...); err != nil {
		return err
	}

	return nil
}

func (w *Watcher) WatchFS(fsys fs.FS, patterns []string) error {
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

	if err := w.Watch(filenames...); err != nil {
		return err
	}
	return nil
}

func (w *Watcher) Watch(filenames ...string) error {
	for _, filename := range filenames {
		if err := w.FSWatcher.Add(filename); err != nil {
			return err
		}
	}

	return nil
}

func (w *Watcher) WatchWorker() {
	for {
		select {
		case event, ok := <-w.FSWatcher.Events:
			if !ok {
				return
			}

			if event.Op&fsnotify.Write == fsnotify.Write {
				w.template.ParseFiles(event.Name)
				log.Println("modified file:", event.Name)
			}
		case err, ok := <-w.FSWatcher.Errors:
			if !ok {
				return
			}

			panic(err)
		}
	}
}
