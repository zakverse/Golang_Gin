package domain

import "time"

type Register struct{
	Username string 		`json:"username"`
	Password string			`json:"password"`
	BirthDate time.Time		`json:"birth_date"`
	CodeReferal string 		`json:"referal_code"`
}

type AuthRepository interface{
	CreateRegister(register Register) error
}