package repository

import (
	"Golang_Gin/internal/domain"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type barangRepository struct {
	db *sql.DB
}

func NewBarangRepository(db *sql.DB) domain.BarangRepository {
	return &barangRepository{db: db}
}

func (r *barangRepository) GetAll() ([]domain.Barang, error) {
	rows, err := r.db.Query("select id, code_barang, nama_barang, code_type_barang, harga_sewa from barang")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	barangs := []domain.Barang{}

	for rows.Next() {
		single := domain.Barang{}
		err := rows.Scan(&single.ID, &single.CodeBarang, &single.NamaBarang, &single.CodeTypeBarang, &single.HargaSewa)
		if err != nil {
			return nil, err
		}
		barangs = append(barangs, single)
	}
	return barangs, nil
}

func (r *barangRepository) GetByID(id uint) (domain.Barang, error) {
	rows, err := r.db.Query("select id, code_barang, nama_barang, code_type_barang, harga_sewa from barang where id = ?", id)
	if err != nil {
		return domain.Barang{}, err
	}
	defer rows.Close()

	barang := domain.Barang{}
	for rows.Next() {
		err := rows.Scan(&barang.ID, &barang.CodeBarang, &barang.NamaBarang, &barang.CodeTypeBarang, &barang.HargaSewa)
		if err != nil {
			return domain.Barang{}, err
		}
	}
	return barang, nil
}

func (r *barangRepository) Create(barang domain.Barang) error {
	_, err := r.db.Exec("insert into barang (code_barang, nama_barang, code_type_barang, harga_sewa) values (?, ?, ?, ?)", barang.CodeBarang, barang.NamaBarang, barang.CodeTypeBarang, barang.HargaSewa)
	if err != nil {
		return err
	}
	return nil
}

func (r *barangRepository) Update(barang domain.Barang) error {
	_, err := r.db.Exec("update barang set code_barang = ?, nama_barang = ?, code_type_barang = ?, harga_sewa = ? where id = ?", barang.CodeBarang, barang.NamaBarang, barang.CodeTypeBarang, barang.HargaSewa, barang.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *barangRepository) Delete(id uint) error {
	_, err := r.db.Exec("delete from barang where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *barangRepository) GetByIDJoin(id uint) (domain.DetailBarang, error) {
	rows, err := r.db.Query("Select a.id , code_barang, nama_barang, harga_sewa, description from barang as a join type as b on a.code_type_barang = b.code_type where a.id = ?", id)
	if err != nil {
		return domain.DetailBarang{}, err
	}
	defer rows.Close()

	barang := domain.DetailBarang{}
	for rows.Next() {
		err := rows.Scan(&barang.IDBarang, &barang.CodeBarang, &barang.NamaBarang, &barang.HargaSewa, &barang.TypeDescription)
		if err != nil {
			return domain.DetailBarang{}, err
		}
	}
	return barang, nil
}

func (r *barangRepository) JoinDetailHargaSewa() ([]domain.DetailHargaSewaBarang, error) {
	rows, err := r.db.Query("select a.id,  a.id_sewa, nama_barang, description, frekuensi_barang ,harga_sewa_barang from detail_sewa as a join barang as c on a.code_barang = c.code_barang join type as b on b.code_type = c.code_type_barang")
	if err != nil {
		return []domain.DetailHargaSewaBarang{}, err
	}
	defer rows.Close()

	barang := []domain.DetailHargaSewaBarang{}
	for rows.Next() {
		single := domain.DetailHargaSewaBarang{}
		err := rows.Scan(&single.ID_Detail_Sewa, &single.IDSewa, &single.Nama_Barang, &single.Description, &single.FrekuensiBarang, &single.HargaSewaBarang)
		if err != nil {
			return nil, err
		}
		barang = append(barang, single)
	}
	return barang, nil
}