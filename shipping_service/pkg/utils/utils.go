package utils

import (
	"encoding/json"
	"errors"

	"github.com/MrRique15/go_rest_api/shipping_service/pkg/shared/models"
	"golang.org/x/crypto/bcrypt"
)

func ByteToInterface(data []byte) models.KafkaOrderEvent {
	var event models.KafkaOrderEvent
	json.Unmarshal(data, &event)
	return event
}

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

func OrderToJson(order models.KafkaOrderEvent) (string, error) {
	jsonOrder, error := json.Marshal(order)

	if error != nil {
		return "", error
	}

	return string(jsonOrder), nil
}