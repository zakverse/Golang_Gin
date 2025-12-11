package domain

type Type struct {
    ID          uint      `json:"id"`
    CodeType    int       `json:"code_type"`
    Description string    `json:"description"`
}

type TypeRepository interface {
	GetAll() ([]Type, error)
	GetByID(id uint) (Type, error)
	// Create(type Type) error
	// Update(type Type) error
	// Delete(id uint) error
}

