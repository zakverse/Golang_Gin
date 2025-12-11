package domain

type Type struct {
	ID          uint   `json:"id"`
	CodeType    int    `json:"code_type"`
	Description string `json:"description"`
}

type TypeRepository interface {
	GetAll() ([]Type, error)
	GetByID(id uint) (Type, error)
	Create(typ Type) error
	Update(typ Type) error
	Delete(id uint) error
}
