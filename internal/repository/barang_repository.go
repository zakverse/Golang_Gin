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
