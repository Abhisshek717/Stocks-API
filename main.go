package main

import (
	"fmt"
	"api/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", jsonContentTypeMiddleware(r)))
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
