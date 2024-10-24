package render

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/Dharmadurai95/go-basic-web-application/pkg/config"
)

var app *config.AppConfig

func NewTemplage(a *config.AppConfig) {
	app = a
}

func ParseTemplate(w http.ResponseWriter, path string) error {
	ts := make(map[string]*template.Template)
	if app.UseCache {
		//create template cache
		ts = app.TemplateCache

	} else {
		ts, _ = CreateTemplateCache()
	}

	//get requested template from cache
	getCacheVal, ok := ts[path]
	if !ok {
		log.Fatal("could not get temaplate from template cache ")
	}

	err := getCacheVal.Execute(w, nil)
	if err != nil {
		return err
	}
	return nil

	// buf := new(bytes.Buffer)
	// err = getCacheVal.Execute(buf, nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// //render the template
	// _, err = buf.WriteTo(w)
	// if err != nil {
	// 	fmt.Println(err)
	// }

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	//get all the files names *.page.html from 'template' dir
	pages, err := filepath.Glob("../../templates/*.page.html")
	if err != nil {
		return cache, err
	}
	fmt.Println(pages)
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		matches, err := filepath.Glob("../../templates/layout.page.html")
		if err != nil {
			return cache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../../templates/layout.page.html")
			if err != nil {
				return cache, err
			}
		}

		cache[name] = ts
	}
	return cache, nil
}
