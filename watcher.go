package temple

import (
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	*fsnotify.Watcher
	t *template.Template
}

func newWatcher(t *template.Template) *Watcher {
	w := Watcher{}

	fsWatcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}

	w.t = t
	w.Watcher = fsWatcher

	go w.WatchWorker()

	return &w
}

func (w *Watcher) WatchWorker() {
	for {
		select {
		case event, ok := <-w.Events:
			if !ok {
				return
			}

			if event.Op&fsnotify.Write == fsnotify.Write {
				name := filepath.Base(event.Name)
				tmpl, err := template.ParseFiles(event.Name)

				if err != nil {
					fmt.Println("failed to parse files", err)
					continue
				}

				*tmpl.Lookup(name) = *tmpl
			}

		case err, ok := <-w.Errors:
			if !ok {
				return
			}

			panic(err)
		}
	}
}
