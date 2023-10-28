package entity

import "clothingShop/db"

type User struct {
	Uuid        uint32 `json:"uuid" gorm:"primaryKey"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	BonusPoints uint32 `json:"bonusPoints"`
	BirthDate   string `json:"birthDate"`
	Login       string `json:"login"`
	Password    string `json:"password"`
	Access      uint8  `json:"access"`
}

func (u *User) TableName() string {
	return "user"
}

func MigrateUser() {
	err := db.DB().AutoMigrate(User{})
	if err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateUser)
}
