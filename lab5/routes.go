package main

import "net/http"


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