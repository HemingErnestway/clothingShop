package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	m := http.NewServeMux()

	m.Handle("/", http.HandlerFunc(mainHandle))

	server := &http.Server{
		Addr:         ":8090",
		Handler:      m,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Listening port :8090...")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
