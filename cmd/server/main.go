package main

import (
	"mable-to-facebook/internal/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/send/facebook", handlers.SendToFacebookHandler)
	port := "8080"
	http.ListenAndServe(":"+port, nil)
}
