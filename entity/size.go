package entity

import "clothingShop/db"

type Size struct {
	Uuid uint32
	Name string
}

func (s *Size) TableName() string {
	return "size"
}

func MigrateSize() {
	err := db.DB().AutoMigrate(Size{})
	if err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateSize)
}
