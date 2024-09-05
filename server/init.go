package server

import (
	"encoding/json"
	"log"
	"net/http"
)


func StartServer(port int) {
	// server a public directory from ../ to /

	http.Handle("/", http.FileServer(http.Dir("../public")))

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		game := CreateGame(1)
		json.NewEncoder(w).Encode(game)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}