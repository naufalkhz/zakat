package models

import (
	"gorm.io/gorm"
)

type Infaq struct {
	gorm.Model
	Judul         string `json:"judul"`
	Deskripsi     string `json:"deskripsi" gorm:"type:text"`
	Gambar        string `json:"gambar"`
	DanaTerkumpul int64  `json:"dana_terkumpul"`
}

type InfaqRequest struct {
	Judul     string `json:"judul" binding:"required"`
	Deskripsi string `json:"deskripsi" binding:"required"`
	Gambar    string `json:"gambar" binding:"required"`
}

type InfaqRiwayat struct {
	gorm.Model
	KodeRiwayat string `json:"kode_riwayat"`
	Nominal     int64  `json:"nominal"`
	Catatan     string `json:"catatan" gorm:"type:text"`
	HambaAllah  bool   `json:"hamba_allah"`

	IdInfaq  uint   `json:"id_infaq"`
	Judul    string `json:"judul"`
	IdUser   uint   `json:"id_user"`
	NamaUser string `json:"nama_user"`
	Email    string `json:"email"`
	PembayaranBank
}

type InfaqRiwayatRequest struct {
	Nominal    int64  `json:"nominal" binding:"required"`
	Catatan    string `json:"catatan"`
	HambaAllah bool   `json:"hamba_allah"`
	IdInfaq    uint   `json:"id_infaq" binding:"required"`
	IdBank     uint   `json:"id_bank" binding:"required"`
}
