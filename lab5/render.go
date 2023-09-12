package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//go:embed templates
var TemplateFS embed.FS

func (a *App) RenderTemplate(w http.ResponseWriter, page string) {

	var t *template.Template
	var err error

	_, exists := a.Cache[page]

	if !exists || a.Config.Env == "dev" {
		t, err = template.ParseFS(
			TemplateFS,
			"templates/"+page+".page.html",
			"templates/navbar.layout.html",
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
