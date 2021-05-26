package main

import (
	"flag"
	"log"
	"net/http"
)

var dest = flag.String("destination", "", "the base of the destination to redirect to")
var addr = flag.String("listen", "0.0.0.0:8080", "port to listen on")

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, *dest+r.URL.Path, http.StatusPermanentRedirect)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("simple-redirect: " + *dest))
}

func main() {
	flag.Parse()
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/", redirect)
	log.Printf("Listening on %s...\n", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatalf("Serving failed: %v.\n", err)
	}
}
