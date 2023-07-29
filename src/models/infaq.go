package models

import (
	"gorm.io/gorm"
)

type Infaq struct {
	gorm.Model
	Judul         string  `json:"judul"`
	Deskripsi     string  `json:"deskripsi" gorm:"type:text"`
	Gambar        string  `json:"gambar"`
	DanaTerkumpul float64 `json:"dana_terkumpul"`
}

type InfaqRequest struct {
	Judul     string `json:"judul" binding:"required"`
	Deskripsi string `json:"deskripsi" binding:"required"`
	Gambar    string `json:"gambar" binding:"required"`
}
