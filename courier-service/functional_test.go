//go:build integration
// +build integration

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// FUNCTIONAL TESTS - Courier Service
// Test yang melibatkan full HTTP request/response
// dan bisa mengakses database jika tersedia

func setupCourierServer() *httptest.Server {
	service := NewCourierService()
	handler := NewCourierHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/delivery", handler.StartDelivery)
	mux.HandleFunc("/courier/deliveries", handler.GetCourierDeliveries)
	mux.HandleFunc("/health", handler.Health)

	return httptest.NewServer(mux)
}

func TestFunctional_CourierStartDelivery(t *testing.T) {
	server := setupCourierServer()
	defer server.Close()

	request := DeliveryRequest{
		Resi:         "RES001",
		CourierID:    1,
		AssignedZone: "Jakarta",
	}

	body, _ := json.Marshal(request)

	resp, err := http.Post(server.URL+"/delivery", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestFunctional_CourierStartDeliveryInvalidRequest(t *testing.T) {
	server := setupCourierServer()
	defer server.Close()

	// Missing required fields
	request := DeliveryRequest{
		Resi:      "",
		CourierID: 0,
	}

	body, _ := json.Marshal(request)

	resp, err := http.Post(server.URL+"/delivery", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", resp.StatusCode)
	}
}

func TestFunctional_CourierGetDeliveries(t *testing.T) {
	server := setupCourierServer()
	defer server.Close()

	resp, err := http.Get(server.URL + "/courier/deliveries?courier_id=1")
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	if result["courier_id"] != float64(1) {
		t.Errorf("expected courier_id 1, got %v", result["courier_id"])
	}
}

func TestFunctional_CourierHealthEndpoint(t *testing.T) {
	server := setupCourierServer()
	defer server.Close()

	resp, err := http.Get(server.URL + "/health")
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}

	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)

	if result["status"] != "healthy" {
		t.Errorf("expected 'healthy', got '%s'", result["status"])
	}
}

func TestFunctional_CourierEndToEnd(t *testing.T) {
	server := setupCourierServer()
	defer server.Close()

	// 1. Start a delivery
	delivery := DeliveryRequest{
		Resi:         "RES-E2E-001",
		CourierID:    5,
		AssignedZone: "Surabaya",
	}

	body, _ := json.Marshal(delivery)
	resp, err := http.Post(server.URL+"/delivery", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to start delivery: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}

	// 2. Get deliveries for courier
	resp, err = http.Get(server.URL + "/courier/deliveries?courier_id=5")
	if err != nil {
		t.Fatalf("failed to get deliveries: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}

	// 3. Check health
	resp, err = http.Get(server.URL + "/health")
	if err != nil {
		t.Fatalf("failed to check health: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}
