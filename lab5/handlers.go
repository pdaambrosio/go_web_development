package main

import "net/http"

// The `HomeHandler` function is a method of the `App` struct. It takes two parameters: `w` of type
// `http.ResponseWriter` and `r` of type `http.Request`.
func (a *App) HomeHandler(w http.ResponseWriter, r *http.Request) {
	a.RenderTemplate(w, "index")
}

// The `ContactHandler` function is a method of the `App` struct. It takes two parameters: `w` of type
// `http.ResponseWriter` and `r` of type `http.Request`.
func (a *App) ContactHandler(w http.ResponseWriter, r *http.Request) {
	a.RenderTemplate(w, "contact")
}

// The `AboutHandler` function is a method of the `App` struct. It takes two parameters: `w` of type
// `http.ResponseWriter` and `r` of type `http.Request`.
func (a *App) AboutHandler(w http.ResponseWriter, r *http.Request) {
	a.RenderTemplate(w, "about")
}
