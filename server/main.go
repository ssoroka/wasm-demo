package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	s := &http.Server{
		Addr:           ":8080",
		Handler:        http.FileServer(http.Dir("client/public")),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
