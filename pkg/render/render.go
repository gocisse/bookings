package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gocisse/bookings/pkg/config"
	"github.com/gocisse/bookings/pkg/models"
)

// Default data to use in all templates
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

//Template function holder to use later
var functions = template.FuncMap{}

//NewTemplate will create a template to use in cache
var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

// this function render template from disk
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	//declare template cache
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// fetch each template and check if it exist
	t, ok := tc[tmpl]
	if !ok {
		fmt.Println("unable to fetch template ")
	}
	//create a buffer to hold our template cache
	buf := new(bytes.Buffer)

	// default data for all templates
	td = AddDefaultData(td)

	// execute template and add it to the empty buffer
	err := t.Execute(buf, td)
	if err != nil {
		log.Fatal(err)
	}

	// Write the buffer to the responseWritter
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("unable to write to the responseWritter ", err)
	}

}

// CreateTemplateCache creates a template cache not from disk
func CreateTemplateCache() (map[string]*template.Template, error) {

	// create a cache to hold the template cachec
	myCache := map[string]*template.Template{}

	// parse pages to put in the template
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	//access each page we parse
	for _, page := range pages {

		// name each page
		name := filepath.Base(page)

		fmt.Println("current page is :", page)
		// create our template sets
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		// Find a match and adding it to the cache
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		// check if find a match in the template folder and parse it globally
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, err
}
