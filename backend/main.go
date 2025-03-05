package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler pour la route "/"
func handler(w http.ResponseWriter, r *http.Request) {
	// Répond avec un message "Hello, world!"
	fmt.Fprintf(w, "test 1")
}

func main() {
	// Définir la route "/"
	http.HandleFunc("/", handler)

	// Lancer le serveur HTTP sur le port 8080
	fmt.Println("Serveur démarré sur le port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
