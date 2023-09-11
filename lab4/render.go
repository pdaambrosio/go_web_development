package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// The `func (a *App) RenderTemplate(w http.ResponseWriter, page string)` function is a method of the
// `App` struct. It is responsible for rendering a template and sending the rendered HTML to the
// `http.ResponseWriter` provided.
func (a *App) RenderTemplate(w http.ResponseWriter, page string) {

	var t *template.Template
	var err error

	_, exists := a.Cache[page]

	if !exists || a.Config.Env == "dev" {
		t, err = template.ParseFiles(
			"templates/"+page+".page.html",
			"templates/base.layout.html",
		)
		if err != nil {
			log.Println(err)
			return
		}
		a.Cache[page] = t
	} else {
		fmt.Println("Cache HIT")
		t = a.Cache[page]
	}

	t.Execute(w, nil)
}
