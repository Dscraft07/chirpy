package main

import (
	"log"
	"net/http"
)

func main() {
	// Vytvoříme nový ServeMux, který bude sloužit jako multiplexer pro HTTP požadavky.
	mux := http.NewServeMux()

	// Vytvoříme nový http.Server a nastavíme:
	// - Addr na ":8080", což znamená, že server bude naslouchat na všech IP adresách na portu 8080.
	// - Handler na náš ServeMux.
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Server starting on :8080...")
	// Spustíme server. Protože v muxu nejsou definovány žádné cesty,
	// server vždy vrátí výchozí odpověď 404 Not Found.
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
