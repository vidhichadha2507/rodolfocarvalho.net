package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", "localhost:8081", "host:port")
	flag.Parse()
	log.Printf("Serving at http://%s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, http.FileServer(http.Dir("."))))
}
