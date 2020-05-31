package db

import "golang.org/x/crypto/bcrypt"

func EncriptPassword(password string) (string, error) {
	salt := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(bytes), err
}