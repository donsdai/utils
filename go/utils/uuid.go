package utils

import "github.com/gofrs/uuid"

func GenerateUUID() (string, error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	return uuid.String(), nil
}
