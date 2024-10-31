package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

var jwtKey []byte

type Claims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
    viper.AddConfigPath("./src/configs")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	jwtKey = []byte(viper.GetString("jwt.key"))
}

func init() {
	initConfig()
}

func GenerateJWT(id uint) (string, error) {
	expirationTime := time.Now().Add(10 * time.Hour)
	claims := &Claims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func VerifyJWT(tokenString string) (uint, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return 0, fmt.Errorf("Invalid token signature")
		}
		return 0, fmt.Errorf("Could not parse token: %v", err)
	}

	if !token.Valid {
		return 0, fmt.Errorf("Token is invalid")
	}

	return claims.ID, nil
}
