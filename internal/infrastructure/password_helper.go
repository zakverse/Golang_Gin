package infrastructure

import (
    "errors"
    "os"

    "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    pepper := os.Getenv("PASSWORD_DZAKI")
    if pepper == "" {
        return "", errors.New("PASSWORD_DZAKI kosong")
    }

    hash, err := bcrypt.GenerateFromPassword(
        []byte(password+pepper),
        bcrypt.DefaultCost,
    )

    return string(hash), err
}

func CheckPassword(hash, password string) error {
    pepper := os.Getenv("PASSWORD_DZAKI")
    if pepper == "" {
        return errors.New("PASSWORD_DZAKI kosong")
    }

    return bcrypt.CompareHashAndPassword(
        []byte(hash),
        []byte(password+pepper),
    )
}