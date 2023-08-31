package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	page, err := os.Open("static" + r.URL.Path + ".html")

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

func main() {
	http.HandleFunc("/", StaticHandler)
	http.ListenAndServe(":8080", nil)
}
