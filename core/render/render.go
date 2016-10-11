package render

import (
	"html/template"
	"log"
)

type (
	assetDirFunc func(name string) ([]string, error)
	assetFunc    func(name string) ([]byte, error)
)

func GetTemplates(assetDir assetDirFunc, asset assetFunc) *template.Template {
	var t, tmpl *template.Template

	// find all tempates
	views, err := assetDir("templates")
	if err != nil {
		log.Fatal("Failed to get views;", err)
	}

	// parse all the templates
	for _, v := range views {
		// initialize t if not done yet
		if t == nil {
			t = template.New(v)
		}
		// get pointer to current template
		if v == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(v)
		}

		// get the contents of the current template
		b, err := asset("templates/" + v)
		if err != nil {
			log.Fatal("Failed to read a template file;", err)
		}

		// parse the template
		_, err = tmpl.Parse(string(b))
		if err != nil {
			log.Fatal("Failed to parse a template;", err)
		}
	}

	return t
}
