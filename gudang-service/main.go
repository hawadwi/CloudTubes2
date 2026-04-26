package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	service := NewSortingService()
	handler := NewSortingHandler(service)

	http.HandleFunc("/sort", handler.StartSort)
	http.HandleFunc("/health", handler.Health)

	fmt.Println("Gudang Service started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}