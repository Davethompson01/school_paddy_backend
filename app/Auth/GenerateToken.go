package auth

import (
	"log"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type Claims struct {
	UserID int `json:"user_id"`

	jwt.RegisteredClaims
}

func secretKey() string {
	godotenv.Load(".env")

	loadSecret := os.Getenv("JWT_SECRET_KEY")
	// []byte loadSecret
	if loadSecret == "" {
		log.Fatal("Failed to Load Port")
	}
	return loadSecret
}

func GenerateToken(userID int) (string, error) {

	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey())
}

func ValidateToken(tokenString string) (*Claims, error) {

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return "", nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}
