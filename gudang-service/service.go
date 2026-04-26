package main

// import (
// 	"errors"
// 	"fmt"
// 	"time"
// )

type SortingService struct {
	// Bisa inject database atau repository di sini nanti
}

func NewSortingService() *SortingService {
	return &SortingService{}
}

// StartSorting - Mulai proses sorting package
// TODO: Implementasi di Tahap 3
func (s *SortingService) StartSorting(pkg *Package) error {
	// TODO: Validasi package tidak nil
	// TODO: Validasi resi tidak kosong
	// TODO: Validasi status = pending
	// TODO: Validasi warehouse_zone tidak kosong
	// TODO: Update status ke "sorting"
	// TODO: Simpan ke database
	return nil
}

// CompleteSorting - Selesaikan proses sorting
// TODO: Implementasi di Tahap 3
func (s *SortingService) CompleteSorting(pkg *Package) error {
	// TODO: Validasi package tidak nil
	// TODO: Validasi status = sorting
	// TODO: Update status ke "ready"
	// TODO: Set SortedAt timestamp
	// TODO: Simpan ke database
	// TODO: Set SortedAt timestamp
	// TODO: Update status ke "ready"
	// TODO: Simpan ke database
	return nil
}

// GetPendingPackages - Ambil semua package dengan status pending
// TODO: Implementasi di Tahap 3
func (s *SortingService) GetPendingPackages(packages []Package) []Package {
	// TODO: Query dari database dengan status = pending
	// TODO: Return hasil packages
	return []Package{}
}

// ValidatePackage - Validasi package data
// TODO: Implementasi di Tahap 3
func (s *SortingService) ValidatePackage(pkg *Package) error {
	// TODO: Validasi package tidak nil
	// TODO: Validasi resi tidak kosong
	// TODO: Validasi user_id > 0
	// TODO: Validasi berat > 0
	// TODO: Validasi warehouse_zone tidak kosong
	return nil
}
