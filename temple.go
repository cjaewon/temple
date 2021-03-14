package temple

import (
	"html/template"
)

type Config struct {
	Name string
	Hot  bool
}

// Template is template.Template wrapper which cover all of html/template method
type Template struct {
	*template.Template
	cfg *Config
}

// New creates a Template instance
func New(c *Config) *Template {
	t := Template{
		Template: template.New(c.Name),
		cfg:      c,
	}

	return &t
}
