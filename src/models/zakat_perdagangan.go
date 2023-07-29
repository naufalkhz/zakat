package models

import (
	"gorm.io/gorm"
)

type ZakatPerdagangan struct {
	gorm.Model
	KodeRiwayat string `json:"kode_riwayat" gorm:"unique"`

	Modal      int64   `json:"modal"`
	Keuntungan int64   `json:"keuntungan"`
	Piutang    int64   `json:"piutang"`
	Utang      int64   `json:"utang"`
	Kerugian   int64   `json:"kerugian"`
	Bayar      float64 `json:"bayar"`
	TransaksiInfo
}

type ZakatPerdaganganRequest struct {
	Modal      int64 `json:"modal" binding:"required"`
	Keuntungan int64 `json:"keuntungan" binding:"required"`
	Piutang    int64 `json:"piutang"`
	Utang      int64 `json:"utang"`
	Kerugian   int64 `json:"kerugian"`
	IdBank     uint  `json:"id_bank" binding:"required"`
}
