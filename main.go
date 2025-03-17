package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}


func main() {
	const port = "8080"

	mux := http.NewServeMux()

	// Přidání readiness endpointu
	mux.HandleFunc("/healthz", healthHandler)

	// Nastavení fileserveru na /app/
	fileServer := http.FileServer(http.Dir("."))
	mux.Handle("/app/", http.StripPrefix("/app", fileServer))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Server listening on port %s\n", port)
	log.Fatal(srv.ListenAndServe())
}

