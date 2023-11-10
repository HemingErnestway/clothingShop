package entity

import "clothingShop/db"

type Country struct {
	Uuid uint32 `json:"uuid" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (c *Country) TableName() string {
	return "country"
}

func MigrateCountry() {
	err := db.DB().AutoMigrate(Country{})
	if err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateCountry)
}
