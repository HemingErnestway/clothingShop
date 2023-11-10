package storage

import (
	"clothingShop/db"
	"clothingShop/entity"
)

func CategoryCreate(category entity.Category) *entity.Category {
	db.DB().Create(&category)
	return &category
}

func CategoryRead(id uint32) *entity.Category {
	var category entity.Category
	db.DB().Table(category.TableName()).Where(
		"uuid = ?", id).Find(&category)
	return &category
}

func CategoriesRead() []*entity.Category {
	var categories []*entity.Category
	db.DB().Find(&categories)
	return categories
}

func CategoryUpdate(new entity.Category, id uint32) *entity.Category {
	var current entity.Category
	db.DB().Table(current.TableName()).Where(
		"uuid = ?", id).Find(&current)

	if new.Name != "" {
		current.Name = new.Name
	}

	db.DB().Save(&current)
	return CategoryRead(id)
}

func CategoryDelete(id uint32) string {
	var category entity.Category
	db.DB().Table(category.TableName()).Where(
		"uuid = ?", id).Delete(&category)
	return "successfully deleted"
}
