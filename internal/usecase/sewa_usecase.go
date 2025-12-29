package usecase

import "Golang_Gin/internal/domain"

type SewaUseCase struct {
	repo domain.SewaRepository
}

func NewSewaUseCase(r domain.SewaRepository) *SewaUseCase {
	return &SewaUseCase{repo: r}
}

func (u *SewaUseCase) Create(sewaBody domain.SewaBody) error {
	return u.repo.Create(sewaBody)
}