package repository

import (
	"database/sql"
	"Golang_Gin/internal/domain"
	_ "github.com/go-sql-driver/mysql"
)

type typeRepository struct {
	db *sql.DB
}

func NewTypeRepository(db *sql.DB) domain.TypeRepository {
	return &typeRepository{db: db}
}

func (r *typeRepository) GetAll() ([]domain.Type, error) {
	rows , err := r.db.Query("select id, code_type, description from type")
	if err != nil{
		return nil, err
	}

	defer rows.Close()

	types := []domain.Type{}

	for rows.Next(){
		single := domain.Type{}
		err := rows.Scan(&single.ID, &single.CodeType, &single.Description)
		if err != nil{
			return nil, err
		}
		types = append(types, single)
	}
	return types, nil
}

