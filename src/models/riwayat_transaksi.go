package models

type RiwayatPembayaranResponse struct {
	ZakatPenghasilan []*ZakatPenghasilan `json:"zakat_penghasilan"`
	ZakatTabungan    []*ZakatTabungan    `json:"zakat_tabungan"`
	ZakatPerdagangan []*ZakatPerdagangan `json:"zakat_perdagangan"`
	ZakatEmas        []*ZakatEmas        `json:"zakat_emas"`
	InfaqRiwayat     []*InfaqRiwayat     `json:"infaq_riwayat"`
}
