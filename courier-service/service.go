package main

// import (
// 	"errors"
// 	"fmt"
// 	"time"
// )

type CourierService struct {
	// Bisa inject database atau repository di sini nanti
}

func NewCourierService() *CourierService {
	return &CourierService{}
}

// StartDelivery - Mulai proses delivery
// TODO: Implementasi di Tahap 3
func (s *CourierService) StartDelivery(delivery *Delivery) error {
	// TODO: Validasi delivery tidak nil
	// TODO: Validasi resi tidak kosong
	// TODO: Validasi status = pending
	// TODO: Validasi courier_id > 0
	// TODO: Update status ke "in_delivery"
	// TODO: Simpan ke database
	return nil
}

// CompleteDelivery - Selesaikan proses delivery
// TODO: Implementasi di Tahap 3
func (s *CourierService) CompleteDelivery(delivery *Delivery) error {
	// TODO: Validasi delivery tidak nil
	// TODO: Validasi status = in_delivery
	// TODO: Update status ke "delivered"
	// TODO: Set DeliveredAt timestamp
	// TODO: Simpan ke database
	return nil
}

// GetCourierDeliveries - Ambil semua delivery untuk courier tertentu
// TODO: Implementasi di Tahap 3
func (s *CourierService) GetCourierDeliveries(deliveries []Delivery, courierID int) []Delivery {
	// TODO: Query dari database berdasarkan courier_id
	// TODO: Return hasil deliveries
	return []Delivery{}
}

// ValidateDelivery - Validasi delivery data
// TODO: Implementasi di Tahap 3
func (s *CourierService) ValidateDelivery(delivery *Delivery) error {
	// TODO: Validasi delivery tidak nil
	// TODO: Validasi resi tidak kosong
	// TODO: Validasi courier_id > 0
	// TODO: Validasi nama_penerima tidak kosong
	// TODO: Validasi alamat_penerima tidak kosong
	return nil
}
