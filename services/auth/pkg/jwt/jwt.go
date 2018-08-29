package jwt

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

var (
	SigningMethod = jwt.SigningMethodHS256
	SigningKey    = []byte(os.Getenv("JWT_SIGNING_KEY"))
)

type Claims jwt.StandardClaims

func NewWithClaims(claims *Claims) (string, error) {
	token := jwt.NewWithClaims(SigningMethod, jwt.StandardClaims(*claims))
	return token.SignedString(SigningKey)
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return SigningKey, nil
}

func ValidateToken(tokenString string) error {
	_, err := jwt.Parse(tokenString, keyFunc)
	return err
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, keyFunc)
	if err != nil {
		return nil, err
	}
	claims := Claims(*token.Claims.(*jwt.StandardClaims))
	return &claims, nil
}
