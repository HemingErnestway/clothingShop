package storage

import (
	"clothingShop/db"
	"clothingShop/entity"
)

func ColorCreate(color entity.Color) *entity.Color {
	db.DB().Create(&color)
	return &color
}

func ColorRead(id uint32) *entity.Color {
	var color entity.Color
	db.DB().Table(color.TableName()).Where(
		"uuid = ?", id).Find(&color)
	return &color
}

func ColorsRead() []*entity.Color {
	var colors []*entity.Color
	db.DB().Find(&colors)
	return colors
}

func ColorUpdate(new entity.Color, id uint32) *entity.Color {
	var current entity.Color
	db.DB().Table(current.TableName()).Where(
		"uuid = ?", id).Find(&current)

	if new.Name != "" {
		current.Name = new.Name
	}

	db.DB().Save(&current)
	return ColorRead(id)
}

func ColorDelete(id uint32) string {
	var color entity.Color
	db.DB().Table(color.TableName()).Where(
		"uuid = ?", id).Delete(&color)
	return "successfully deleted"
}
