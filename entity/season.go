package entity

import "clothingShop/db"

type Season struct {
	Uuid uint32 `json:"uuid" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (s *Season) TableName() string {
	return "season"
}

func MigrateSeason() {
	err := db.DB().AutoMigrate(Season{})
	if err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateSeason)
}
