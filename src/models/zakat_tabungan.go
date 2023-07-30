package models

import (
	"gorm.io/gorm"
)

type ZakatTabungan struct {
	gorm.Model
	KodeRiwayat   string  `json:"kode_riwayat" gorm:"unique"`
	SaldoTabungan int64   `json:"saldo_tabungan"`
	Bunga         int64   `json:"bunga"`
	Bayar         float64 `json:"bayar"`
	PembayaranInfo
}

type ZakatTabunganRequest struct {
	SaldoTabungan int64 `json:"saldo_tabungan" binding:"required"`
	Bunga         int64 `json:"bunga"`
	IdBank        uint  `json:"id_bank" binding:"required"`
}
