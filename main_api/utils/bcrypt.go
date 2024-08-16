package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashFromPassword(password string) (string, error){
	pass := []byte(password)

    hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
    if err != nil {
        return "", errors.New("error during password hashing")
    }

    return string(hash), nil
}

func ComparePasswords(password string, hashedPass string) error {
	pass := []byte(password)
	hashedPassByte := []byte(hashedPass)

	err := bcrypt.CompareHashAndPassword(hashedPassByte, pass)

	if err != nil{
		return err
	}
	
	return nil
}