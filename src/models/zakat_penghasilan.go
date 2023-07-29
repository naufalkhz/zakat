package models

import (
	"gorm.io/gorm"
)

type ZakatPenghasilan struct {
	gorm.Model
	KodeRiwayat          string  `json:"kode_riwayat" gorm:"unique"`
	Penghasilan          int64   `json:"penghasilan"`
	PendapatanLain       int64   `json:"pendapatan_lain"`
	PengeluaranKebutuhan int64   `json:"pengeluaran_kebutuhan"`
	JenisPenghasilan     string  `json:"jenis_penghasilan"`
	Bayar                float64 `json:"bayar"`
	TransaksiInfo
}

type ZakatPenghasilanRequest struct {
	Penghasilan          int64  `json:"penghasilan" binding:"required"`
	PendapatanLain       int64  `json:"pendapatan_lain"`
	PengeluaranKebutuhan int64  `json:"pengeluaran_kebutuhan"`
	JenisPenghasilan     string `json:"jenis_penghasilan" binding:"required"`
	IdBank               uint   `json:"id_bank" binding:"required"`
}
