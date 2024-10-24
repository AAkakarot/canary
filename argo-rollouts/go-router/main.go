// main.go
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("OK")); err != nil {
			log.Printf("Error writing response: %v", err)
		}
	})

	// Add a dummy endpoint
	http.HandleFunc("/dummy", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("Dummy endpoint works")); err != nil {
			log.Printf("Error writing response: %v", err)
		}
	})

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
