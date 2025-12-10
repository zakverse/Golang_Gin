package models

type Barang struct {
	ID             int     `json:"id"`
	CodeBarang     string  `json:"code_barang" binding:"required"`
	NamaBarang     string  `json:"nama_barang" binding:"required"`
	CodeTypeBarang uint    `json:"code_type_barang"` // Foreign Key ke Type
	HargaSewa      float64 `json:"harga_sewa" binding:"required"`
	IsActive       bool    `json:"is_active"`

	// Relasi
	Type Type 				`json:"type"`
}

/*INSERT INTO barang (code_barang, nama_barang, code_type_barang, harga_sewa, isActive) VALUES
-- Type 1: Televisi
('TV001', 'Televisi LED 32 Inch', 1, 50000, 1),
('TV002', 'Televisi LED 43 Inch', 1, 70000, 1),
('TV003', 'Televisi Smart TV 55 Inch', 1, 120000, 1),

-- Type 2: Komputer
('PC001', 'PC Office Core i3', 2, 80000, 1),
('PC002', 'PC Gaming RTX 3060', 2, 150000, 1),
('PC003', 'PC Editing Core i7', 2, 130000, 1),

-- Type 3: Peralatan Game
('GM001', 'PlayStation 4', 3, 90000, 1),
('GM002', 'PlayStation 5', 3, 150000, 1),
('GM003', 'Nintendo Switch', 3, 80000, 1),

-- Type 4: Camera
('CM001', 'Kamera DSLR Canon', 4, 100000, 1),
('CM002', 'Kamera Mirrorless Sony', 4, 130000, 1),
('CM003', 'Action Cam GoPro', 4, 60000, 1),

-- Type 5: Kipas Angin
('KP001', 'Kipas Angin Berdiri', 5, 20000, 1),
('KP002', 'Kipas Angin Dinding', 5, 15000, 1),
('KP003', 'Kipas Angin Meja', 5, 10000, 1),

-- Type 6: Setrika
('ST001', 'Setrika Listrik Standard', 6, 10000, 1),
('ST002', 'Setrika Uap Portable', 6, 15000, 1),
('ST003', 'Setrika Uap Besar', 6, 20000, 1),

-- Type 7: Rice Cooker
('RC001', 'Rice Cooker Mini 0.6L', 7, 15000, 1),
('RC002', 'Rice Cooker 1.8L', 7, 20000, 1),
('RC003', 'Rice Cooker Digital 2L', 7, 30000, 1);*/

var DataBarang = []Barang{
	{ID: 1, CodeBarang: "TV001", NamaBarang: "Televisi LED 32 Inch", CodeTypeBarang: 1, HargaSewa: 50000, IsActive: true},
	{ID: 2, CodeBarang: "TV002", NamaBarang: "Televisi LED 43 Inch", CodeTypeBarang: 1, HargaSewa: 70000, IsActive: true},
	{ID: 3, CodeBarang: "TV003", NamaBarang: "Televisi Smart TV 55 Inch", CodeTypeBarang: 1, HargaSewa: 120000, IsActive: true},

	{ID: 4, CodeBarang: "PC001", NamaBarang: "PC Office Core i3", CodeTypeBarang: 2, HargaSewa: 80000, IsActive: true},
	{ID: 5, CodeBarang: "PC002", NamaBarang: "PC Gaming RTX 3060", CodeTypeBarang: 2, HargaSewa: 150000, IsActive: true},
	{ID: 6, CodeBarang: "PC003", NamaBarang: "PC Editing Core i7", CodeTypeBarang: 2, HargaSewa: 130000, IsActive: true},

	{ID: 7, CodeBarang: "GM001", NamaBarang: "PlayStation 4", CodeTypeBarang: 3, HargaSewa: 90000, IsActive: true},
	{ID: 8, CodeBarang: "GM002", NamaBarang: "PlayStation 5", CodeTypeBarang: 3, HargaSewa: 150000, IsActive: true},
	{ID: 9, CodeBarang: "GM003", NamaBarang: "Nintendo Switch", CodeTypeBarang: 3, HargaSewa: 80000, IsActive: true},

	{ID: 10, CodeBarang: "CM001", NamaBarang: "Kamera DSLR Canon", CodeTypeBarang: 4, HargaSewa: 100000, IsActive: true},
	{ID: 11, CodeBarang: "CM002", NamaBarang: "Kamera Mirrorless Sony", CodeTypeBarang: 4, HargaSewa: 130000, IsActive: true},
	{ID: 12, CodeBarang: "CM003", NamaBarang: "Action Cam GoPro", CodeTypeBarang: 4, HargaSewa: 60000, IsActive: true},

	{ID: 13, CodeBarang: "KP001", NamaBarang: "Kipas Angin Berdiri", CodeTypeBarang: 5, HargaSewa: 20000, IsActive: true},
	{ID: 14, CodeBarang: "KP002", NamaBarang: "Kipas Angin Dinding", CodeTypeBarang: 5, HargaSewa: 15000, IsActive: true},
	{ID: 15, CodeBarang: "KP003", NamaBarang: "Kipas Angin Meja", CodeTypeBarang: 5, HargaSewa: 10000, IsActive: true},
	{ID: 16, CodeBarang: "ST001", NamaBarang: "Setrika Listrik Standard", CodeTypeBarang: 6, HargaSewa: 10000, IsActive: true},

	{ID: 17, CodeBarang: "ST002", NamaBarang: "Setrika Uap Portable", CodeTypeBarang: 6, HargaSewa: 15000, IsActive: true},
	{ID: 18, CodeBarang: "ST003", NamaBarang: "Setrika Uap Besar", CodeTypeBarang: 6, HargaSewa: 20000, IsActive: true},

	{ID: 19, CodeBarang: "RC001", NamaBarang: "Rice Cooker Mini 0.6L", CodeTypeBarang: 7, HargaSewa: 15000, IsActive: true},
	{ID: 20, CodeBarang: "RC002", NamaBarang: "Rice Cooker 1.8L", CodeTypeBarang: 7, HargaSewa: 20000, IsActive: true},
	{ID: 21, CodeBarang: "RC003", NamaBarang: "Rice Cooker Digital 2L", CodeTypeBarang: 7, HargaSewa: 30000, IsActive: true},
}

