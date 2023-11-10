package entity

import "clothingShop/db"

type Brand struct {
	Uuid uint32 `json:"uuid" gorm:"primaryKay"`
	Name string `json:"name"`
}

func (u *Brand) TableName() string {
	return "brand"
}

func MigrateBrand() {
	err := db.DB().AutoMigrate(Brand{})
	if err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateBrand)
}
