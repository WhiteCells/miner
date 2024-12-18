package utils

import (
	"github.com/google/uuid"
)

func GenerateInviteCode() string {
	uuid := uuid.New()
	return uuid.String()
}
