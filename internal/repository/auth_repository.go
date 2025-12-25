package repository

import (
	"Golang_Gin/internal/domain"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"time"
)

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) domain.AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) CreateRegister(register domain.Register) error {
	query := "insert into user (username , password, join_date, birth_date, referal_code )values(? , ? , ?, ?, ?)"
	_, err := r.db.Exec(query, register.Username, register.Password, time.Now(), register.BirthDate, register.CodeReferal)
	if err != nil {
		return err
	}
	return nil
}
