package main

import (
	//handlers "Golang_Gin/Handlers"
	"Golang_Gin/internal/delivery/http"
	"Golang_Gin/internal/infrastructure"
	"Golang_Gin/internal/repository"
	"Golang_Gin/internal/usecase"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	db, err := infrastructure.NewMySQL()

	if err != nil {
		panic(err)
	}

	renv := godotenv.Load("../../.env")
	if renv != nil {
		log.Fatal("Gagal load .env")
	}

	if os.Getenv("PASSWORD_DZAKI") == "" {
		log.Fatal("PASSWORD_DZAKI belum diset")
	}

	repoTypes := repository.NewTypeRepository(db)
	ucTypes := usecase.NewTypeUseCase(repoTypes)
	http.NewTypeHandler(r, ucTypes)

	repoBarang := repository.NewBarangRepository(db)
	ucBarang := usecase.NewBarangUseCase(repoBarang)
	http.NewBarangHandler(r, ucBarang)

	repoAuth := repository.NewAuthRepository(db)
	ucAuth := usecase.NewAuthUseCase(repoAuth)
	http.NewAuthHandler(r, ucAuth)

	r.Run(":8080")
}
