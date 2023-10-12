package entity

import (
	"time"
)

type User struct {
	Uuid        uint32
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
