package entity

import "clothingShop/db"

type Product struct {
	Uuid            uint32  `json:"uuid" gorm:"primaryKey"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	NewPrice        float64 `json:"newPrice"`
	DiscountPercent uint8   `json:"discountPercent"`
	Quantity        uint32  `json:"quantity"`
	CategoryId      uint32  `json:"categoryId"`
	SeasonId        uint32  `json:"seasonId"`
	ColorId         uint32  `json:"colorId"`
	CountryId       uint32  `json:"countryId"`
	GenderId        uint32  `json:"genderId"`
	AgeGroupId      uint32  `json:"ageGroupId"`
	BrandId         uint32  `json:"brandId"`
	SizeId          uint32  `json:"sizeId"`
}

func (p *Product) TableName() string {
	return "product"
}

func MigrateProduct() {
	err := db.DB().AutoMigrate(Product{})
	if err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateProduct)
}
