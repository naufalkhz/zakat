package models

import "gorm.io/gorm"

type Bank struct {
	gorm.Model
	NamaBank   string `json:"nama_bank" binding:"required"`
	NoRekening string `json:"no_rekening" binding:"required,numeric"`
	AtasNama   string `json:"atas_nama" binding:"required"`
}
