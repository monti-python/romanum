package main

import (
	"fmt"
	"log"
	"net/http"
	"romanum/internal/handlers"
)



func main() {
	http.HandleFunc("/convert", handlers.ConversionHandler)
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
