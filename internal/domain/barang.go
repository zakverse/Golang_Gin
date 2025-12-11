package domain

type Barang struct {
	ID             int     `json:"id"`
	CodeBarang     string  `json:"code_barang" binding:"required"`
	NamaBarang     string  `json:"nama_barang" binding:"required"`
	CodeTypeBarang uint    `json:"code_type_barang"` // Foreign Key ke Type
	HargaSewa      float64 `json:"harga_sewa" binding:"required"`
	IsActive       bool    `json:"is_active"`

	// Relasi
	Type Type `json:"type"`
}

type BarangRepository interface {
	GetAll() ([]Barang, error)
	GetByID(id uint) (Barang, error)
	Create(barang Barang) error
	Update(barang Barang) error
	Delete(id uint) error
}
