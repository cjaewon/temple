package temple

import "io/fs"

func (t *Template) ParseFiles(filenames ...string) (*Template, error) {
	tmpl, err := t.Template.ParseFiles(filenames...)
	if err != nil {
		return nil, err
	}

	if t.cfg.Hot {
		watch(filenames...)
	}

	t.Template = tmpl
	return t, nil
}

func (t *Template) ParseGlob(pattern string) (*Template, error) {
	tmpl, err := t.Template.ParseGlob(pattern)
	if err != nil {
		return nil, err
	}

	if t.cfg.Hot {
		if err := watchGlob(pattern); err != nil {
			return nil, err
		}
	}

	t.Template = tmpl
	return t, nil
}

func (t *Template) ParseFS(fs fs.FS, patterns ...string) (*Template, error) {
	tmpl, err := t.Template.ParseFS(fs, patterns...)
	if err != nil {
		return nil, err
	}

	if t.cfg.Hot {
		if err := watchFS(fs, patterns); err != nil {
			return nil, err
		}
	}

	t.Template = tmpl
	return t, nil
}
