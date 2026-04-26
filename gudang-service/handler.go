// package main

// import (
// 	"encoding/json"
// 	"net/http"
// )

// type SortingHandler struct {
// 	service *SortingService
// }

// func NewSortingHandler(service *SortingService) *SortingHandler {
// 	return &SortingHandler{service: service}
// }

// // POST /sort
// func (h *SortingHandler) StartSort(w http.ResponseWriter, r *http.Request) {
// 	var req SortRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	// Validate
// 	if req.Resi == "" {
// 		http.Error(w, "Resi is required", http.StatusBadRequest)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]string{"status": "sorting started"})
// }

// // GET /health
// func (h *SortingHandler) Health(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
// }

package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// Interface untuk memudahkan mocking di test
type SortingServiceInterface interface {
	StartSorting(pkg *Package) error
	CompleteSorting(pkg *Package) error
	GetPendingPackages(packages []Package) []Package
	ValidatePackage(pkg *Package) error
}

type SortingHandler struct {
	service SortingServiceInterface
}

func NewSortingHandler(service SortingServiceInterface) *SortingHandler {
	return &SortingHandler{service: service}
}

// POST /sort
func (h *SortingHandler) StartSort(w http.ResponseWriter, r *http.Request) {
	var req SortRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validasi manual (karena gak pakai framework seperti gin)
	if req.Resi == "" {
		http.Error(w, "Resi is required", http.StatusBadRequest)
		return
	}

	if req.WarehouseZone == "" {
		http.Error(w, "Warehouse zone is required", http.StatusBadRequest)
		return
	}

	if req.Status == "" {
		http.Error(w, "Status is required", http.StatusBadRequest)
		return
	}

	// Validasi status (biar realistis)
	if req.Status != "sorting" {
		http.Error(w, "Invalid status, must be 'sorting'", http.StatusBadRequest)
		return
	}

	// Simulasi update waktu sorting
	now := time.Now()

	// (Optional) panggil service kalau ada
	// h.service.StartSort(req.Resi, req.WarehouseZone, req.Status, now)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":          "sorting started",
		"resi":            req.Resi,
		"warehouse_zone":  req.WarehouseZone,
		"sorted_at":       now,
	})
}

// GET /health
func (h *SortingHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{
		"status": "healthy",
	})
}
