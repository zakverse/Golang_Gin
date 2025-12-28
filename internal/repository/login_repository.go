package repository

import (
	"Golang_Gin/internal/domain"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type loginRepository struct {
	db *sql.DB
}

func NewLoginRepository(db *sql.DB) domain.LoginUserRepository {
	return &loginRepository{db: db}
}

func (l *loginRepository) Login(login domain.Login) (domain.LoginResponse, error) {
	query := "select id, username, password from user where username = ?"
	rows, err := l.db.Query(query, login.Username)

	if err != nil {
		return domain.LoginResponse{}, err
	}
	defer rows.Close()

	var user domain.LoginResponse
	for rows.Next() {
		err := rows.Scan(&user.UserID, &user.Username, &user.Password)
		if err != nil {
			return domain.LoginResponse{}, err
		}
	}
	return user, nil
}
