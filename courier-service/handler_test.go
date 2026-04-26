package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

// =====================
// TEST START DELIVERY
// =====================
func TestStartDelivery_Success(t *testing.T) {
	service := NewCourierService()
	handler := NewCourierHandler(service)

	body := []byte(`{
		"resi": "RESI123",
		"courier_id": 1,
		"assigned_zone": "JKT"
	}`)

	req := httptest.NewRequest(http.MethodPost, "/delivery", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	handler.StartDelivery(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", res.StatusCode)
	}
}

func TestStartDelivery_InvalidJSON(t *testing.T) {
	service := NewCourierService()
	handler := NewCourierHandler(service)

	body := []byte(`{invalid-json}`)

	req := httptest.NewRequest(http.MethodPost, "/delivery", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.StartDelivery(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", res.StatusCode)
	}
}

func TestStartDelivery_MissingField(t *testing.T) {
	service := NewCourierService()
	handler := NewCourierHandler(service)

	body := []byte(`{
		"resi": "",
		"courier_id": 0,
		"assigned_zone": ""
	}`)

	req := httptest.NewRequest(http.MethodPost, "/delivery", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.StartDelivery(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", res.StatusCode)
	}
}

// =====================
// TEST GET DELIVERY
// =====================
func TestGetCourierDeliveries_Success(t *testing.T) {
	service := NewCourierService()
	handler := NewCourierHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/courier/deliveries?courier_id=1", nil)
	w := httptest.NewRecorder()

	handler.GetCourierDeliveries(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", res.StatusCode)
	}
}

func TestGetCourierDeliveries_InvalidID(t *testing.T) {
	service := NewCourierService()
	handler := NewCourierHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/courier/deliveries?courier_id=abc", nil)
	w := httptest.NewRecorder()

	handler.GetCourierDeliveries(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", res.StatusCode)
	}
}

func TestGetCourierDeliveries_MissingID(t *testing.T) {
	service := NewCourierService()
	handler := NewCourierHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/courier/deliveries", nil)
	w := httptest.NewRecorder()

	handler.GetCourierDeliveries(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", res.StatusCode)
	}
}

// =====================
// TEST HEALTH
// =====================
func TestHealth(t *testing.T) {
	service := NewCourierService()
	handler := NewCourierHandler(service)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	handler.Health(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", res.StatusCode)
	}
}
