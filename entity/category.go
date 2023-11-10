package entity

import "clothingShop/db"

type Category struct {
	Uuid uint32 `json:"uuid" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (c *Category) TableName() string {
	return "category"
}

func MigrateCategory() {
	err := db.DB().AutoMigrate(Category{})
	if err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateCategory)
}
