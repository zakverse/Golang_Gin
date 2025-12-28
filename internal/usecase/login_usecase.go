package usecase

import (
	"Golang_Gin/internal/domain"
	"Golang_Gin/internal/infrastructure"
	"errors"
)

type LoginUseCase struct {
	repo domain.LoginUserRepository
}

func NewLoginUseCase(r domain.LoginUserRepository) *LoginUseCase {
	return &LoginUseCase{repo: r}
}

func (l *LoginUseCase) Login(login domain.Login) (domain.LoginResponse, error) {
	if login.Username == "" || login.Password == "" {
		return domain.LoginResponse{}, errors.New("Username or Password is empty")
	}
	var value, err = l.repo.Login(login)
	if err != nil {
		return domain.LoginResponse{}, err
	}
	if value.UserID == 0 {
		return domain.LoginResponse{}, errors.New("username atau password salah")
	}
	var erorPassword = infrastructure.CheckPassword(value.Password, login.Password)
	if erorPassword != nil {
		return domain.LoginResponse{}, errors.New("username atau password salah")
	}
	value.Password = ""
	return value, err
}
