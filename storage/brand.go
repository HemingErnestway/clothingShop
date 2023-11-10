package storage

import (
	"clothingShop/db"
	"clothingShop/entity"
)

func BrandCreate(brand entity.Brand) *entity.Brand {
	db.DB().Create(&brand)
	return &brand
}

func BrandRead(id uint32) *entity.Brand {
	var brand entity.Brand
	db.DB().Table(brand.TableName()).Where(
		"uuid = ?", id).Find(&brand)
	return &brand
}

func BrandsRead() []*entity.Brand {
	var brands []*entity.Brand
	db.DB().Find(&brands)
	return brands
}

func BrandUpdate(new entity.Brand, id uint32) *entity.Brand {
	var current entity.Brand
	db.DB().Table(current.TableName()).Where(
		"uuid = ?", id).Find(&current)

	if new.Name != "" {
		current.Name = new.Name
	}

	db.DB().Save(&current)
	return BrandRead(id)
}

func BrandDelete(id uint32) string {
	var brand entity.Brand
	db.DB().Table(brand.TableName()).Where(
		"uuid = ?", id).Delete(&brand)
	return "successfully deleted"
}
