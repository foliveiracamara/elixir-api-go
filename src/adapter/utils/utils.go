package utils

import (
	"time"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	uuid := uuid.New()
	return uuid.String()
}

func GetCurrentDateTime()	string {
	return time.Now().Format("01-02-2006 15:04")
}

func GetCurrentDate() string {
	return time.Now().Format("01-02-2006")
}