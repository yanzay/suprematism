package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("p", 8080, "Port to listen")
	flag.Parse()
	addr := fmt.Sprintf("localhost:%d", *port)
	log.Printf("Listening on http://%s", addr)
	http.ListenAndServe(addr, http.FileServer(http.Dir(".")))
}
