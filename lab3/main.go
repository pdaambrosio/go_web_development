package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// The line `var env = "dev"` is declaring a variable named `env` and assigning it the value "dev".
// This variable is used to determine the environment in which the application is running.
var env = "dev"
var cache map[string]*template.Template

// The HomeHandler function renders the "index" template for the home page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index")
}

// The ContactHandler function renders a template named "contact" in response to an HTTP request.
func ContactHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "contact")
}

// The RenderTemplate function renders a template using the provided http.ResponseWriter and page name,
// and caches the template if it doesn't already exist or if the environment is set to "dev".
func RenderTemplate(w http.ResponseWriter, page string) {

	var t *template.Template
	var err error

	_, exists := cache[page]

	if !exists || env == "dev" {
		t, err = template.ParseFiles(
			"templates/"+page+".page.html",
			"templates/base.layout.html",
		)
		if err != nil {
			log.Println(err)
			return
		}
		cache[page] = t
	} else {
		fmt.Println("Cache HIT")
		t = cache[page]
	}

	t.Execute(w, nil)
}

// The main function sets up the routing for different HTTP requests and starts the server on port
// 3000.
func main() {
	cache = make(map[string]*template.Template)

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/contact", ContactHandler)

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, "about")
	})

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", nil)
}
