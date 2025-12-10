package main

import (
	//handlers "Golang_Gin/Handlers"
	"Golang_Gin/internal/infrastructure"
	"Golang_Gin/internal/repository"
	"Golang_Gin/internal/usecase"
	"Golang_Gin/internal/delivery/http"
	"github.com/gin-gonic/gin"
)
func main() {
	r := gin.Default();
	db, err := infrastructure.NewMySQL();

	if err != nil{
		panic(err);
	}

	repoTypes := repository.NewTypeRepository(db)
	ucTypes := usecase.NewTypeUseCase(repoTypes)

	http.NewTypeHandler(r, ucTypes)
	r.Run(":8081")
}