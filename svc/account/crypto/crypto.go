package crypt

import "golang.org/x/crypto/bcrypt"

type Crypt interface {
	HashPassword(string) (string, error)
	CheckHashedPassword(string, string) bool
}

type BCrypt struct {
	Crypt
}

func NewBCrypt() *BCrypt {
	return &BCrypt{}
}

func (c *BCrypt) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (c *BCrypt) CheckHashedPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
