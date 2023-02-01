package core

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Create the JWT key used to create the signature
var JwtKey = []byte("d2f00e8a-72ae-11ea-8103-f01898ebdc2e")

// jwt.StandardClaims is an embedded type
type Claims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

// GenerateToken generates a new token
func GenerateToken(id string) string {
	// Declare the expiration time of the token
	// expirationTime := time.Now().Add(72 * time.Hour)
	expirationTime := time.Now().AddDate(1, 0, 0)
	claims := &Claims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	tokenString, _ := token.SignedString(JwtKey)
	return tokenString
}

func VerifyToken(tokenBody string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(
		tokenBody,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		},
	)
	if err != nil {
		return nil, errors.New("malformed token")
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
