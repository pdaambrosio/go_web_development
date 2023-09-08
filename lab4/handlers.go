package main

import "net/http"

func (a *App) HomeHandler(w http.ResponseWriter, r *http.Request) {
	a.RenderTemplate(w, "index")
}

func (a *App) ContactHandler(w http.ResponseWriter, r *http.Request) {
	a.RenderTemplate(w, "contact")
}

func (a *App) AboutHandler(w http.ResponseWriter, r *http.Request) {
	a.RenderTemplate(w, "about")
}
