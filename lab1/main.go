package main

import (
	"log"
	"net/http"
	"os"
)

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	page, err := os.ReadFile("static/index.html")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
		log.Fatal(err)
	}

	w.Write(page)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	page, err := os.ReadFile("static/about.html")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
		log.Fatal(err)
	}

	w.Write(page)
}

func main() {
	http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/about", AboutHandler)
	http.ListenAndServe(":8080", nil)
}
