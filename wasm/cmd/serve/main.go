package main

import (
	"log"
	"net/http"
)

const (
	host = ":8080"
	dir  = "static"
)

func main() {
	fs := http.FileServer(http.Dir(dir))
	log.Printf("serving %s on http://localhost%s", dir, host)
	http.ListenAndServe(host, fs)
}
