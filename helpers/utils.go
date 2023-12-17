package helpers

import (
	"github.com/google/uuid"
)

type Utils struct {
}

func (u *Utils) GenerateRandomID() string {
	id := uuid.New()
	return id.String()
}
