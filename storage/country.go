package storage

import (
	"clothingShop/db"
	"clothingShop/entity"
)

func CountryCreate(country entity.Country) *entity.Country {
	db.DB().Create(&country)
	return &country
}

func CountryRead(id uint32) *entity.Country {
	var country entity.Country
	db.DB().Table(country.TableName()).Where(
		"uuid = ?", id).Find(&country)
	return &country
}

func CountriesRead() []*entity.Country {
	var categories []*entity.Country
	db.DB().Find(&categories)
	return categories
}

func CountryUpdate(new entity.Country, id uint32) *entity.Country {
	var current entity.Country
	db.DB().Table(current.TableName()).Where(
		"uuid = ?", id).Find(&current)

	if new.Name != "" {
		current.Name = new.Name
	}

	db.DB().Save(&current)
	return CountryRead(id)
}

func CountryDelete(id uint32) string {
	var country entity.Country
	db.DB().Table(country.TableName()).Where(
		"uuid = ?", id).Delete(&country)
	return "successfully deleted"
}
