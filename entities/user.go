package entities

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Uuid        uuid.UUID
	Name        string
	Surname     string
	Email       string
	Address     string
	BonusPoints uint
	BirthDate   time.Time
	Login       string
	Password    string
	Access      uint8
}
