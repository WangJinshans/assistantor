package utils

import "github.com/google/uuid"

func GetUuidString() string {
	return uuid.New().String()
}

