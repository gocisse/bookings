package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

//This package will hold all the configs
//to use accross all packages

type AppConfig struct {
	UseCache bool
	TemplateCache map[string]*template.Template
	InProduction bool 
	Session *scs.SessionManager
}
