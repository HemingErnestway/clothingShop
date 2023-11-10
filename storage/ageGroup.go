package storage

import (
	"clothingShop/db"
	"clothingShop/entity"
)

func AgeGroupCreate(ageGroup entity.AgeGroup) *entity.AgeGroup {
	db.DB().Create(&ageGroup)
	return &ageGroup
}

func AgeGroupRead(id uint32) *entity.AgeGroup {
	var ageGroup entity.AgeGroup
	db.DB().Table(ageGroup.TableName()).Where(
		"uuid = ?", id).Find(&ageGroup)
	return &ageGroup
}

func AgeGroupsRead() []*entity.AgeGroup {
	var ageGroups []*entity.AgeGroup
	db.DB().Find(&ageGroups)
	return ageGroups
}

func AgeGroupUpdate(new entity.AgeGroup, id uint32) *entity.AgeGroup {
	var current entity.AgeGroup
	db.DB().Table(current.TableName()).Where(
		"uuid = ?", id).Find(&current)

	if new.Name != "" {
		current.Name = new.Name
	}

	db.DB().Save(&current)
	return AgeGroupRead(id)
}

func AgeGroupDelete(id uint32) string {
	var ageGroup entity.AgeGroup
	db.DB().Table(ageGroup.TableName()).Where(
		"uuid = ?", id).Delete(&ageGroup)
	return "successfully deleted"
}
