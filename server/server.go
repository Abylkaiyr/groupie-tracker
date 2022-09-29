package server

import (
	"fmt"
	"log"
	"net/http"
)

func Server() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("ui"))
	mux.Handle("/ui/", http.StripPrefix("/ui", fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/detail", details)
	fmt.Println("Starting the web server on http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
		return
	}
}
