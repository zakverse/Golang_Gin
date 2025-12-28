package domain

type Login struct {
	Username string 	`json:"username"`
	Password string 	`json:"password"`
}

type LoginResponse struct {
	UserID int 			`json:"id"`
	Username string 	`json:"username"`
	Password string 	`json:"password"`
}

type LoginUserRepository interface {
	Login(login Login) (LoginResponse, error)
}