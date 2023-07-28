package models

import (
	"gorm.io/gorm"
)

type ZakatPenghasilan struct {
	gorm.Model
	KodeRiwayat          string  `json:"kodeRiwayat"`
	Penghasilan          int64   `json:"penghasilan"`
	PendapatanLain       int64   `json:"pendapatanLain"`
	PengeluaranKebutuhan int64   `json:"pengeluaranKebutuhan"`
	JenisPenghasilan     string  `json:"jenisPenghasilan"`
	HargaEmas            int64   `json:"hargaEmas"`
	Bayar                float64 `json:"bayar"`
	IdUser               uint    `json:"idUser"`
	EmailUser            string  `json:"emailUser"`
}
