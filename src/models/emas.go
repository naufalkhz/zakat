package models

import (
	"gorm.io/gorm"
)

type Emas struct {
	gorm.Model
	Harga  int64  `json:"harga"`
	Tipe   string `json:"tipe"`
	Sumber string `json:"sumber"`
}
