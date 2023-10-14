package entity

import (
	"time"
)

type User struct {
	Uuid        uint32    `json:"uuid"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	Email       string    `json:"email"`
	Address     string    `json:"address"`
	BonusPoints uint      `json:"bonusPoints"`
	BirthDate   time.Time `json:"birthDate"`
	Login       string    `json:"login"`
	Password    string    `json:"password"`
	Access      uint8     `json:"access"`
}
