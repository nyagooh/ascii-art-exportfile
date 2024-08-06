package main

import (
	"fmt"
	"net/http"
	"os"

	handlers "asciiart/handlers"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("<Usage>", "<go run .> or <go run main.go>")
		os.Exit(1)
	}
	http.HandleFunc("/", handlers.Request)
	http.HandleFunc("/ascii-art", handlers.Post)
	http.HandleFunc("/error", handlers.ErrorHandler)
	http.HandleFunc("/export",handlers.ErrorHandler)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Server is starting on http://localhost:8060")
	err := http.ListenAndServe(":8060", nil)
	if err != nil {
		fmt.Println("Error starting server")
		os.Exit(1)
	}
}
