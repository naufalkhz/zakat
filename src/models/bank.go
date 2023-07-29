package models

import "gorm.io/gorm"

type Bank struct {
	gorm.Model
	Nama       string `json:"nama" binding:"required"`
	NoRekening string `json:"no_rekening" binding:"required,numeric"`
	AtasNama   string `json:"atas_nama" binding:"required"`
}
