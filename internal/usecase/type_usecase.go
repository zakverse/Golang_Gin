package usecase

import(
	"Golang_Gin/internal/domain"
)


type TypeUseCase struct {
	repo domain.TypeRepository
}

func NewTypeUseCase(r domain.TypeRepository) *TypeUseCase{
	return &TypeUseCase{repo: r}
}

//=================================================================

func (u *TypeUseCase) GetAll() ([]domain.Type, error){
	return u.repo.GetAll()
}

func (u *TypeUseCase) GetByID(id uint) (domain.Type, error){
	return u.repo.GetByID(id)
}


//=================================================================
