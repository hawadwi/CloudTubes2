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

// ============================================
// FUNCTIONAL TESTS - Gudang Service
// Test yang melibatkan full HTTP request/response
// dan bisa mengakses database jika tersedia
// ============================================

func setupGudangServer() *httptest.Server {
	service := NewSortingService()
	handler := NewSortingHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/sort", handler.StartSort)
	mux.HandleFunc("/health", handler.Health)

	return httptest.NewServer(mux)
}

func TestFunctional_StartSort(t *testing.T) {
	server := setupGudangServer()
	defer server.Close()

	request := SortRequest{
		Resi:          "RES001",
		WarehouseZone: "Jakarta",
		Status:        "sorting",
	}

	body, _ := json.Marshal(request)

	resp, err := http.Post(server.URL+"/sort", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	if result["resi"] != "RES001" {
		t.Errorf("expected resi 'RES001', got %v", result["resi"])
	}
}

func TestFunctional_StartSortInvalidRequest(t *testing.T) {
	server := setupGudangServer()
	defer server.Close()

	request := SortRequest{
		Resi:          "",
		WarehouseZone: "",
		Status:        "",
	}

	body, _ := json.Marshal(request)

	resp, err := http.Post(server.URL+"/sort", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", resp.StatusCode)
	}
}

func TestFunctional_Health(t *testing.T) {
	server := setupGudangServer()
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

func TestFunctional_InvalidJSON(t *testing.T) {
	server := setupGudangServer()
	defer server.Close()

	resp, err := http.Post(server.URL+"/sort", "application/json",
		bytes.NewReader([]byte("invalid json")))
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", resp.StatusCode)
	}
}

func TestFunctional_SortEndToEnd(t *testing.T) {
	server := setupGudangServer()
	defer server.Close()

	// 1. Start sorting
	request := SortRequest{
		Resi:          "RES-E2E-001",
		WarehouseZone: "Surabaya",
		Status:        "sorting",
	}

	body, _ := json.Marshal(request)

	resp, err := http.Post(server.URL+"/sort", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to start sort: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}

	// 2. Check health
	resp, err = http.Get(server.URL + "/health")
	if err != nil {
		t.Fatalf("failed to check health: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}
