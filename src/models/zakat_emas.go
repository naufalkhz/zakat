package models

import (
	"gorm.io/gorm"
)

type ZakatEmas struct {
	gorm.Model
	KodeRiwayat string  `json:"kode_riwayat" gorm:"unique"`
	Emas        int64   `json:"emas"`
	Bayar       float64 `json:"bayar"`
	TransaksiInfo
}

type ZakatEmasRequest struct {
	Emas   int64 `json:"emas" binding:"required"`
	IdBank uint  `json:"id_bank" binding:"required"`
}
