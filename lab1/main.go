package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// The StaticHandler function serves static files such as HTML and CSS files.
// It takes a http.ResponseWriter and a http.Request as arguments.
// It opens the requested file and copies it to the http.ResponseWriter.
// If the file does not exist, it returns a 404 Not Found error.
func StaticHandler(w http.ResponseWriter, r *http.Request) {
	page, err := os.Open("static" + r.URL.Path)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
		fmt.Println(err)
		return
	}

	if strings.HasSuffix(r.URL.Path, ".css") {
		w.Header().Add("Content-Type", "text/css")
	}

	io.Copy(w, page)
}

// The main function sets up a static file server and listens for incoming HTTP requests on port 8080.
func main() {
	http.HandleFunc("/", StaticHandler)
	http.ListenAndServe(":8080", nil)
}
