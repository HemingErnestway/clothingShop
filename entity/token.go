package entity

import (
	"clothingShop/db"
	"time"
)

type Token struct {
	Tid     uint32    `json:"-" gorm:"primaryKey"`
	Uid     uint32    `json:"uid"`
	Token   string    `json:"token"`
	Expired time.Time `json:"expired"`
}

func (_ Token) TableName() string {
	return "token"
}

func MigrateToken() {
	err := db.DB().AutoMigrate(Token{})
	if err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateToken)
}
