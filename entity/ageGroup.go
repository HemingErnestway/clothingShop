package entity

import "clothingShop/db"

type AgeGroup struct {
	Uuid uint32 `json:"uuid" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (a *AgeGroup) TableName() string {
	return "ageGroup"
}

func MigrateAgeGroup() {
	err := db.DB().AutoMigrate(AgeGroup{})
	if err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateAgeGroup)
}
