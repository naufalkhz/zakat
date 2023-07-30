package models

type TransaksiResponse struct {
	KodeRiwayat string  `json:"kode_riwayat"`
	Bayar       float64 `json:"bayar"`
	Bank        Bank    `json:"bank"`
}

type TransaksiInfo struct {
	HargaEmas int64  `json:"harga_emas"`
	IdUser    uint   `json:"id_user"`
	NamaUser  string `json:"nama_user"`
	EmailUser string `json:"email_user"`
	TransaksiBank
}

type TransaksiBank struct {
	IdBank     uint   `json:"id_bank"`
	Nama       string `json:"nama_bank"`
	NoRekening string `json:"no_rekening"`
	AtasNama   string `json:"atas_nama"`
}
