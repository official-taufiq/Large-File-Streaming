package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	mux.HandleFunc("POST /register", register)
	mux.Handle("POST /upload", authMiddleware(http.HandlerFunc(upload)))

	fmt.Printf("Serving on port%v\n", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
