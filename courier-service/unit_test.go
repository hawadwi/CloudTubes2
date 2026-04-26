package main

import (
	"testing"
)

// UNIT TESTS - CourierService
// Tidak boleh mengakses database atau external service

func TestStartDeliverySuccess(t *testing.T) {
	service := NewCourierService()

	delivery := &Delivery{
		CourierID:      1,
		Resi:           "RES001",
		NamaPenerima:   "John Doe",
		AlamatPenerima: "Jl. Sudirman",
		Status:         "pending",
	}

	err := service.StartDelivery(delivery)
	if err != nil {
		t.Errorf("StartDelivery failed: %v", err)
	}

	if delivery.Status != "in_delivery" {
		t.Errorf("Expected status 'in_delivery', got '%s'", delivery.Status)
	}
}

func TestStartDeliveryInvalidCourier(t *testing.T) {
	service := NewCourierService()

	delivery := &Delivery{
		CourierID: 0,
		Resi:      "RES001",
		Status:    "pending",
	}

	err := service.StartDelivery(delivery)
	if err == nil {
		t.Error("Expected error for invalid courier_id, got nil")
	}
}

func TestStartDeliveryEmptyResi(t *testing.T) {
	service := NewCourierService()

	delivery := &Delivery{
		CourierID: 1,
		Resi:      "",
		Status:    "pending",
	}

	err := service.StartDelivery(delivery)
	if err == nil {
		t.Error("Expected error for empty resi, got nil")
	}
}

func TestStartDeliveryNotPendingStatus(t *testing.T) {
	service := NewCourierService()

	delivery := &Delivery{
		CourierID: 1,
		Resi:      "RES001",
		Status:    "in_delivery",
	}

	err := service.StartDelivery(delivery)
	if err == nil {
		t.Error("Expected error for non-pending delivery, got nil")
	}
}

func TestCompleteDeliverySuccess(t *testing.T) {
	service := NewCourierService()

	delivery := &Delivery{
		Resi:           "RES001",
		Status:         "in_delivery",
		CourierID:      1,
		AlamatPenerima: "Jl. Sudirman",
	}

	err := service.CompleteDelivery(delivery)
	if err != nil {
		t.Errorf("CompleteDelivery failed: %v", err)
	}

	if delivery.Status != "delivered" {
		t.Errorf("Expected status 'delivered', got '%s'", delivery.Status)
	}

	if delivery.DeliveredAt == nil {
		t.Error("DeliveredAt should not be nil")
	}
}

func TestCompleteDeliveryNotInProgress(t *testing.T) {
	service := NewCourierService()

	delivery := &Delivery{
		Resi:      "RES001",
		Status:    "pending",
		CourierID: 1,
	}

	err := service.CompleteDelivery(delivery)
	if err == nil {
		t.Error("Expected error for non-in_delivery status, got nil")
	}
}

func TestCompleteDeliveryNil(t *testing.T) {
	service := NewCourierService()

	err := service.CompleteDelivery(nil)
	if err == nil {
		t.Error("Expected error for nil delivery, got nil")
	}
}

func TestGetCourierDeliveries(t *testing.T) {
	service := NewCourierService()

	deliveries := []Delivery{
		{CourierID: 1, Resi: "RES001", Status: "pending"},
		{CourierID: 2, Resi: "RES002", Status: "pending"},
		{CourierID: 1, Resi: "RES003", Status: "delivered"},
	}

	courierDeliveries := service.GetCourierDeliveries(deliveries, 1)
	if len(courierDeliveries) != 2 {
		t.Errorf("Expected 2 deliveries for courier 1, got %d", len(courierDeliveries))
	}
}

func TestGetCourierDeliveriesNoMatch(t *testing.T) {
	service := NewCourierService()

	deliveries := []Delivery{
		{CourierID: 1, Resi: "RES001", Status: "pending"},
		{CourierID: 2, Resi: "RES002", Status: "pending"},
	}

	courierDeliveries := service.GetCourierDeliveries(deliveries, 999)
	if len(courierDeliveries) != 0 {
		t.Errorf("Expected 0 deliveries for courier 999, got %d", len(courierDeliveries))
	}
}

func TestValidateDeliverySuccess(t *testing.T) {
	service := NewCourierService()

	delivery := &Delivery{
		Resi:           "RES001",
		CourierID:      1,
		NamaPenerima:   "John Doe",
		AlamatPenerima: "Jl. Sudirman",
	}

	err := service.ValidateDelivery(delivery)
	if err != nil {
		t.Errorf("ValidateDelivery failed: %v", err)
	}
}

func TestValidateDeliveryEmptyResi(t *testing.T) {
	service := NewCourierService()

	delivery := &Delivery{
		Resi:           "",
		CourierID:      1,
		NamaPenerima:   "John Doe",
		AlamatPenerima: "Jl. Sudirman",
	}

	err := service.ValidateDelivery(delivery)
	if err == nil {
		t.Error("Expected error for empty resi, got nil")
	}
}

func TestValidateDeliveryInvalidCourier(t *testing.T) {
	service := NewCourierService()

	delivery := &Delivery{
		Resi:           "RES001",
		CourierID:      0,
		NamaPenerima:   "John Doe",
		AlamatPenerima: "Jl. Sudirman",
	}

	err := service.ValidateDelivery(delivery)
	if err == nil {
		t.Error("Expected error for invalid courier_id, got nil")
	}
}

func TestValidateDeliveryMissingReceiver(t *testing.T) {
	service := NewCourierService()

	delivery := &Delivery{
		Resi:           "RES001",
		CourierID:      1,
		NamaPenerima:   "John Doe",
		AlamatPenerima: "",
	}

	err := service.ValidateDelivery(delivery)
	if err == nil {
		t.Error("Expected error for empty alamat_penerima, got nil")
	}
}

func TestValidateDeliveryNil(t *testing.T) {
	service := NewCourierService()

	err := service.ValidateDelivery(nil)
	if err == nil {
		t.Error("Expected error for nil delivery, got nil")
	}
}

// ============================================
// BENCHMARKS
// ============================================

func BenchmarkStartDelivery(b *testing.B) {
	service := NewCourierService()
	delivery := &Delivery{
		CourierID: 1,
		Resi:      "RES001",
		Status:    "pending",
	}

	for i := 0; i < b.N; i++ {
		service.StartDelivery(delivery)
		delivery.Status = "pending"
	}
}

func BenchmarkValidateDelivery(b *testing.B) {
	service := NewCourierService()
	delivery := &Delivery{
		Resi:           "RES001",
		CourierID:      1,
		NamaPenerima:   "John Doe",
		AlamatPenerima: "Jl. Sudirman",
	}

	for i := 0; i < b.N; i++ {
		service.ValidateDelivery(delivery)
	}
}
