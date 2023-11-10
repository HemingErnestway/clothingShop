package storage

import (
	"clothingShop/db"
	"clothingShop/entity"
)

func GenderCreate(gender entity.Gender) *entity.Gender {
	db.DB().Create(&gender)
	return &gender
}

func GenderRead(id uint32) *entity.Gender {
	var gender entity.Gender
	db.DB().Table(gender.TableName()).Where(
		"uuid = ?", id).Find(&gender)
	return &gender
}

func GendersRead() []*entity.Gender {
	var genders []*entity.Gender
	db.DB().Find(&genders)
	return genders
}

func GenderUpdate(new entity.Gender, id uint32) *entity.Gender {
	var current entity.Gender
	db.DB().Table(current.TableName()).Where(
		"uuid = ?", id).Find(&current)

	if new.Name != "" {
		current.Name = new.Name
	}

	db.DB().Save(&current)
	return GenderRead(id)
}

func GenderDelete(id uint32) string {
	var gender entity.Gender
	db.DB().Table(gender.TableName()).Where(
		"uuid = ?", id).Delete(&gender)
	return "successfully deleted"
}
