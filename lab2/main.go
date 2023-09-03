package main

import (
	"net/http"
)

// The main function sets up a static file server to serve files from the "static" directory on port
// 8080.
func main() {
	staticFiles := http.Dir("./static")
	staticHandler := http.FileServer(staticFiles)

	http.Handle("/", staticHandler)
	http.ListenAndServe(":8080", nil)
}
