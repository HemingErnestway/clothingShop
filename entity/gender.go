package entity

import "clothingShop/db"

type Gender struct {
	Uuid uint32 `json:"uuid" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (g *Gender) TableName() string {
	return "gender"
}

func MigrateGender() {
	err := db.DB().AutoMigrate(Gender{})
	if err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateGender)
}
