package main

import "net/http"


// The `Routes()` function is defining the routes for a web application. It creates a new
// `http.ServeMux` object, which is a request multiplexer that matches incoming requests against a list
// of registered patterns and calls the corresponding handler for the pattern that most closely matches
// the request URL.
func (a *App) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", a.HomeHandler)
	mux.HandleFunc("/contact", a.ContactHandler)
	mux.HandleFunc("/about", a.AboutHandler)

	mux.Handle("/static/",
	  http.StripPrefix("/static/",
	    http.FileServer(http.Dir("./static/"))))

	return mux
}