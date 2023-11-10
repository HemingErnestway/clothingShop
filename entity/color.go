package entity

import "clothingShop/db"

type Color struct {
	Uuid uint32 `json:"uuid" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (c *Color) TableName() string {
	return "color"
}

func MigrateColor() {
	err := db.DB().AutoMigrate(Color{})
	if err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateColor)
}
