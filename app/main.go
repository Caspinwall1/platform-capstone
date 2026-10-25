package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

var version = "dev" // overwritten at build time

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		host, _ := os.Hostname()
		writeJSON(w, map[string]string{
			"message": "hello from the Week 1 Hands-On Guide capstone",
			"host":    host,
		})
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{"status": "ok"})
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{"version": version})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("listening on :%s (version %s)", port, version)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}