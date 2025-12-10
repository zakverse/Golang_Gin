package models

type DetailSewa struct {
    ID                 uint           `json:"id"`
    IDSewa             uint           `json:"id_sewa"` // Foreign Key to Sewa
    CodeBarang         string         `json:"code_barang"` // Foreign Key to Barang 
    HargaSewaBarang    float64        `json:"harga_sewa_barang"`
    FrekuensiBarang    int            `json:"frekuensi_barang"`

    // Relasi
    SewaHeader         Sewa           `json:"sewa_header"`
    Barang             Barang         `json:"barang"`
}