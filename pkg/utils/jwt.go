package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte(os.Getenv("SECRET_KEY"))

func GetSecretKey() ([]byte, error) {
	if len(secret) == 0 {
		return nil, errors.New("secret key is not set in .env file")
	}

	return secret, nil
}

type CustomClaims struct {
	Id      int    `json:"id"`
	Nama    string `json:"nama"`
	IsAdmin bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

func GenerateToken(id int, nama string, isAdmin bool) (string, error) {

	secretKey, err := GetSecretKey()
	if err != nil {
		return "", nil
	}

	claims := &CustomClaims{
		Id:      id,
		Nama:    nama,
		IsAdmin: isAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func ParseToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, err
	}

	return claims, nil

}
