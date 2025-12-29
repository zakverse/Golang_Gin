package domain

import "time"

type DetailSewaBody struct {
	CodeBarang      string  `json:"code_barang"`
	HargaSewaBarang float64 `json:"harga_sewa_barang"`
	FrekuensiBarang int     `json:"frekuensi_barang"`
}

type SewaBody struct {
	DetailSewaBody    []DetailSewaBody `json:"detail_sewa_body"`
	IDUser            uint             `json:"id_user"`
	CodeStruk         string           `json:"code_struk"`
	TanggalPeminjaman time.Time        `json:"tanggal_peminjaman"`
	TanggalJatuhTempo time.Time        `json:"tanggal_jatuh_tempo"`
	TanggalPengembalian time.Time      `json:"tanggal_pengembalian"`
	PointReward       int              `json:"point_reward"`
	DendaKeterlambatan float64         `json:"denda_keterlambatan"`
}

type SewaRepository interface {
	Create(sewaBody SewaBody) error
}
