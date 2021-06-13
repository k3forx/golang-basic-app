package config

import (
	"text/template"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool // false is development mode as cache is reloaded every time
	TemplateCache map[string]*template.Template
}
