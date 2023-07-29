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
	HargaEmas            int64   `json:"harga_emas"`
	Bayar                float64 `json:"bayar"`
	IdUser               uint    `json:"id_user"`
	EmailUser            string  `json:"email_user"`
	IdBank               uint    `json:"id_bank"`
	NamaBank             string  `json:"nama_bank"`
	AtasNama             string  `json:"atas_nama"`
	NoRekening           string  `json:"uint"`
	// Bank Bank `json:"bank" gorm:"foreign"`
}

type ZakatPenghasilanRequest struct {
	Penghasilan          int64  `json:"penghasilan" binding:"required"`
	PendapatanLain       int64  `json:"pendapatan_lain" binding:"required"`
	PengeluaranKebutuhan int64  `json:"pengeluaran_kebutuhan" binding:"required"`
	JenisPenghasilan     string `json:"jenis_penghasilan" binding:"required"`
	IdBank               uint   `json:"id_bank" binding:"required"`
}

type ZakatPenghasilanResponse struct {
	KodeRiwayat string  `json:"kode_riwayat"`
	Bayar       float64 `json:"bayar"`
	Bank        Bank    `json:"bank"`
}
