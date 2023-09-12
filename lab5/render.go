package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//go:embed templates
// The line `var TemplateFS embed.FS` is declaring a variable named `TemplateFS` of type `embed.FS`.
var TemplateFS embed.FS

// The `RenderTemplate` function is responsible for rendering a template and sending the rendered HTML
// to the HTTP response writer (`w`).
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

	contact := struct {
		Email string
		Address string
		Phone string
		Twitter string
		Facebook string
	}{
		Email: "mail@mail.com",
		Address: "123 Main St, New York, NY 10030",
		Phone: "555-555-5555",
	}

	t.Execute(w, contact)
}
