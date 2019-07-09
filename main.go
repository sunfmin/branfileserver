package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var address *string = flag.String("a", "localhost:8000", "Address (host:port) to listen on")

func main() {
	fs := http.FileServer(http.Dir("."))

	f := func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		if r.Method == "POST" {
			evt := r.FormValue("__execute_event__")
			if len(evt) > 0 {
				jsonPath := fmt.Sprintf(".%s.%s.json", r.URL.Path, evt)
				log.Println("Response with file: ", jsonPath)
				http.ServeFile(w, r, jsonPath)
				return
			}
		}
		fs.ServeHTTP(w, r)
	}

	log.Printf("serving . on http://%s", *address)
	log.Fatal(http.ListenAndServe(*address, http.HandlerFunc(f)))
}
