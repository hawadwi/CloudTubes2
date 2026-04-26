package main

import "time"

type Package struct {
	UserID         int        `json:"user_id"`
	Resi           string     `json:"resi"`
	NamaBarang     string     `json:"nama_barang"`
	Berat          int        `json:"berat"`
	Dimensi        string     `json:"dimensi"`
	Jenis          string     `json:"jenis"`
	AlamatPengirim string     `json:"alamat_pengirim"`
	AlamatPenerima string     `json:"alamat_penerima"`
	Status         string     `json:"status"` // pending, sorting, ready
	WarehouseZone  string     `json:"warehouse_zone"`
	CreatedAt      time.Time  `json:"created_at"`
	SortedAt       *time.Time `json:"sorted_at"`
}

type SortRequest struct {
	Resi          string `json:"resi" binding:"required"`
	WarehouseZone string `json:"warehouse_zone" binding:"required"`
	Status        string `json:"status" binding:"required"`
}