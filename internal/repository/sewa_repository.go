package repository

import (
	"Golang_Gin/internal/domain"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type sewaRepository struct {
	db *sql.DB
}

func NewSewaRepository(db *sql.DB) domain.SewaRepository {
	return &sewaRepository{db: db}
}

func (s *sewaRepository) Create(sewaBody domain.SewaBody) error {
	tx, errTx := s.db.Begin()
	if errTx != nil {
		return errTx 
	}
	result, err := tx.Exec("insert into sewa (id_user, code_struk, tanggal_peminjaman, tanggal_jatuh_tempo, tanggal_pengembalian, point_reward, denda_keterlambatan) values (?, ?, ?, ?, ?, ?, ?)",

		sewaBody.IDUser,
		sewaBody.CodeStruk,
		sewaBody.TanggalPeminjaman,
		sewaBody.TanggalJatuhTempo,
		sewaBody.TanggalPengembalian,
		sewaBody.PointReward,
		sewaBody.DendaKeterlambatan)

	if err != nil {
		tx.Rollback()
		return err	
	}

	idsewa, errs := result.LastInsertId()
	for _, detail := range sewaBody.DetailSewaBody {
		_, errs = tx.Exec("insert into detail_sewa (id_sewa, code_barang, harga_sewa_barang, frekuensi_barang) values (?, ?, ?, ?)", 
			idsewa, detail.CodeBarang, detail.HargaSewaBarang, detail.FrekuensiBarang)
		if errs != nil {
			tx.Rollback()
			return errs
		}
	}
	tx.Commit()
	return nil
}
