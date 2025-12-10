package models

import "time"

type Sewa struct {
    ID                    uint           `json:"id"`
    IDUser                uint           `json:"id_user"` // Foreign Key to User
    CodeStruk             string         `json:"code_struk"`
    TanggalPeminjaman     time.Time      `json:"tanggal_peminjaman"`
    TanggalJatuhTempo     time.Time      `json:"tanggal_jatuh_tempo"`
    TanggalPengembalian   *time.Time     `json:"tanggal_pengembalian"`
    PointReward           int            `json:"point_reward"`
    DendaKeterlambatan    float64        `json:"denda_keterlambatan"`

    // Relasi
    User                  User           `json:"user"`
}