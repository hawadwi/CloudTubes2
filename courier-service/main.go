package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	service := NewCourierService()
	handler := NewCourierHandler(service)

	http.HandleFunc("/delivery", handler.StartDelivery)
	http.HandleFunc("/courier/deliveries", handler.GetCourierDeliveries)
	http.HandleFunc("/health", handler.Health)

	fmt.Println("Courier Service started on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
