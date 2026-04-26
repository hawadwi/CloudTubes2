package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// Interface untuk memudahkan mocking di test
type CourierServiceInterface interface {
	StartDelivery(delivery *Delivery) error
	CompleteDelivery(delivery *Delivery) error
	GetCourierDeliveries(deliveries []Delivery, courierID int) []Delivery
	ValidateDelivery(delivery *Delivery) error
}

type CourierHandler struct {
	service CourierServiceInterface
}

func NewCourierHandler(service CourierServiceInterface) *CourierHandler {
	return &CourierHandler{service: service}
}

// POST /delivery
func (h *CourierHandler) StartDelivery(w http.ResponseWriter, r *http.Request) {
	var req DeliveryRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Resi == "" || req.CourierID <= 0 || req.AssignedZone == "" {
		http.Error(w, "resi, courier_id, assigned_zone are required", http.StatusBadRequest)
		return
	}

	delivery := &Delivery{
		Resi:         req.Resi,
		CourierID:    req.CourierID,
		AssignedZone: req.AssignedZone,
		Status:       "pending",
		CreatedAt:    time.Now(),
	}

	// panggil service
	if err := h.service.StartDelivery(delivery); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(delivery)
}

// GET /courier/deliveries?courier_id=1
func (h *CourierHandler) GetCourierDeliveries(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("courier_id")

	if idStr == "" {
		http.Error(w, "courier_id is required", http.StatusBadRequest)
		return
	}

	courierID, err := strconv.Atoi(idStr)
	if err != nil || courierID <= 0 {
		http.Error(w, "invalid courier_id", http.StatusBadRequest)
		return
	}

	// dummy data (sementara kalau belum DB)
	all := []Delivery{
		{
			Resi:      "RESI001",
			CourierID: 1,
			Status:    "delivered",
		},
		{
			Resi:      "RESI002",
			CourierID: 2,
			Status:    "in_delivery",
		},
	}

	result := h.service.GetCourierDeliveries(all, courierID)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"courier_id": courierID,
		"count":      len(result),
		"data":       result,
	})
}

// GET /health
func (h *CourierHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "healthy",
	})
}
