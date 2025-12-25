package usecase

import (
	"Golang_Gin/internal/domain"
	"Golang_Gin/internal/infrastructure"
)

type AuthUseCase struct {
	repo domain.AuthRepository
}

func NewAuthUseCase(r domain.AuthRepository) *AuthUseCase {
	return &AuthUseCase{repo: r}
}

func (u *AuthUseCase) Register(register domain.Register) error {

	newPassword , err := infrastructure.HashPassword(register.Password)
	if err != nil{
		return err
	}
	var NewRegister domain.Register = domain.Register{
		Username: register.Username,
		Password: newPassword,
		BirthDate: register.BirthDate,
		CodeReferal: register.CodeReferal,
	}
	
	return u.repo.CreateRegister(NewRegister)
}