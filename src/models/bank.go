package models

import "gorm.io/gorm"

type Bank struct {
	gorm.Model
	Nama       string `json:"name"`
	NoRekening string `json:"no_rekening"`
	AtasNama   string `json:"atas_nama"`
}
