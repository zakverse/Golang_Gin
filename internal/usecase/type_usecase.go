package usecase

import (
	"Golang_Gin/internal/domain"
	"errors"
	"strings"
)

type TypeUseCase struct {
	repo domain.TypeRepository
}

func NewTypeUseCase(r domain.TypeRepository) *TypeUseCase {
	return &TypeUseCase{repo: r}
}

//=================================================================

func (u *TypeUseCase) GetAll() ([]domain.Type, error) {
	return u.repo.GetAll()
}

func (u *TypeUseCase) GetByID(id uint) (domain.Type, error) {
	return u.repo.GetByID(id)
}

func (u *TypeUseCase) Create(typ domain.Type) error {
	if typ.Description == "zaki" {
		return errors.New("Cok")
	}
	return u.repo.Create(typ)
}

func (u *TypeUseCase) Update(typ domain.Type) error {
	kodeBaru := 0
	if typ.CodeType == 1000 {
		kodeBaru = 2000
		typ.CodeType = kodeBaru

	} else if typ.CodeType == 100 {
		kodeBaru = 231
		typ.CodeType = kodeBaru

	} else if typ.CodeType%3 == 1 {
		kodeBaru = typ.CodeType * 3
		typ.CodeType = kodeBaru

	} else if typ.CodeType%2 == 0 {
		kodeBaru = typ.CodeType + 40
		typ.CodeType = kodeBaru
	}

	if typ.Description == "abang" {
		typ.Description = "farras ganteng"
	}
	if len(typ.Description) == 10 {
		kodeBaru = 401
		typ.CodeType = kodeBaru
	}
	if strings.EqualFold(strings.TrimSpace(typ.Description), strings.ToLower("salonpaskoyo")) {
		return errors.New("ErrorKoyo")
	}
	return u.repo.Update(typ)
}

func (u *TypeUseCase) Delete(id uint) error {
	return u.repo.Delete(id)
}

//=================================================================
