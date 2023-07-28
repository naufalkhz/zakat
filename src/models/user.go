package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama         string `json:"nama" binding:"required"`
	Email        string `json:"email" gorm:"unique" binding:"required,email"`
	Password     string `json:"password"`
	JenisKelamin string `json:"jenis_kelamin"`
	TanggalLahir string `json:"tanggal_lahir"`
	NomorTelepon string `json:"nomor_telepon" gorm:"unique" binding:"required,numeric"`
	Kota         string `json:"kota"`
	Pekerjaan    string `json:"pekerjaan"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
