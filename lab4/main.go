package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var env = "dev"
var cache map[string]*template.Template

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index")
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "contact")
}

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

func main() {
	cache = make(map[string]*template.Template)

	config := Config{
		Port:    "8080",
		Env:     "dev",
		Version: "1.0.0",
	}

	// app := App{
	// 	Config: config,
	// 	Cache:  cache,
	// }

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/contact", ContactHandler)

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, "about")
	})

	log.Printf("Server running on port %s Version: %s-%s", config.Port, config.Version, config.Env)
	err := http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		log.Println(err)
	}
}
