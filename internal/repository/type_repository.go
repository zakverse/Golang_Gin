package repository

import (
	"Golang_Gin/internal/domain"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type typeRepository struct {
	db *sql.DB
}

func NewTypeRepository(db *sql.DB) domain.TypeRepository {
	return &typeRepository{db: db}
}

func (r *typeRepository) GetAll() ([]domain.Type, error) {
	rows, err := r.db.Query("select id, code_type, description from type")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	types := []domain.Type{}

	for rows.Next() {
		single := domain.Type{}
		err := rows.Scan(&single.ID, &single.CodeType, &single.Description)
		if err != nil {
			return nil, err
		}
		types = append(types, single)
	}
	return types, nil
}

func (r *typeRepository) GetByID(id uint) (domain.Type, error) {
	rows, err := r.db.Query("select code_type, description from type where id = ?", id)
	if err != nil {
		return domain.Type{}, err
	}
	defer rows.Close()

	typ := domain.Type{}
	for rows.Next() {
		err := rows.Scan(&typ.CodeType, &typ.Description)
		if err != nil {
			return domain.Type{}, err
		}
	}
	return typ, nil
}

func (r *typeRepository) Create(typ domain.Type) error {
	_, err := r.db.Exec("insert into type (code_type, description) values (?, ?)", typ.CodeType, typ.Description)
	if err != nil {
		return err
	}
	return nil
}

func (r *typeRepository) Update(typ domain.Type) error {
	_, err := r.db.Exec("update type set code_type = ?, description = ? where id = ?", typ.CodeType, typ.Description, typ.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *typeRepository) Delete(id uint) error {
	_, err := r.db.Exec("delete from type where id = ?", id)
	if err != nil {
		return err
	}
	return nil
}
