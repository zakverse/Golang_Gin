package usecase

import (
	"Golang_Gin/internal/domain"
)

type BarangUseCase struct {
	repo domain.BarangRepository
}

func NewBarangUseCase(r domain.BarangRepository) *BarangUseCase {
	return &BarangUseCase{repo: r}
}

func (u *BarangUseCase) GetAll() ([]domain.Barang, error) {
	return u.repo.GetAll()
}

func (u *BarangUseCase) GetByID(id uint) (domain.Barang, error) {
	return u.repo.GetByID(id)
}

func (u *BarangUseCase) Create(barang domain.Barang) error {
	return u.repo.Create(barang)
}

func (u *BarangUseCase) Update(barang domain.Barang) error {
	return u.repo.Update(barang)
}

func (u *BarangUseCase) Delete(id uint) error {
	return u.repo.Delete(id)
}	