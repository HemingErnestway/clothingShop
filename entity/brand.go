package entity

import "clothingShop/db"

type Brand struct {
	Uuid uint32 `json:"uuid" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (b *Brand) TableName() string {
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
