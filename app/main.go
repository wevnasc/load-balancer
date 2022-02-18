package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Index(serverName string) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-Type", "text/html")
		fmt.Fprint(rw, serverName)
	}
}

func run() error {
	port := os.Getenv("PORT")
	serverName := os.Getenv("SERVER_NAME")

	http.HandleFunc("/", Index(serverName))

	return http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
