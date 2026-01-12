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
	mux.HandleFunc("GET /files", filesList)
	mux.Handle("GET /download/{fileName}", authMiddleware(http.HandlerFunc(stream)))

	fmt.Printf("Serving on port%v\n", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
