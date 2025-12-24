package domain

type Barang struct {
	ID             uint    `json:"id"`
	CodeBarang     string  `json:"code_barang" binding:"required"`
	NamaBarang     string  `json:"nama_barang" binding:"required"`
	CodeTypeBarang uint    `json:"code_type_barang"` // Foreign Key ke Type
	HargaSewa      float64 `json:"harga_sewa" binding:"required"`
	IsActive       bool    `json:"is_active"`

	// Relasi
	Type Type `json:"type"`
}

type DetailBarang struct {
	IDBarang uint 			    `json:"id"`
	CodeBarang string			`json:"code_barang"`
	NamaBarang string			`json:"nama_barang"`
	HargaSewa float64			`json:"harga_sewa"`
	TypeDescription string		`json:"description"`
}

type DetailHargaSewaBarang struct {
	ID_Detail_Sewa uint			`json:"id"`
	IDSewa uint				    `json:"id_sewa"`
	Nama_Barang string			`json:"nama_barang"`
	Description string			`json:"description"`
	FrekuensiBarang uint  		`json:"frekuensi_barang"`
	HargaSewaBarang float64		`json:"harga_sewa_barang"`
}


type BarangRepository interface {
	GetAll() ([]Barang, error)
	GetByID(id uint) (Barang, error)
	GetByIDJoin(id uint) (DetailBarang, error)
	JoinDetailHargaSewa() ([]DetailHargaSewaBarang, error)
	Create(barang Barang) error
	Update(barang Barang) error
	Delete(id uint) error
}	
