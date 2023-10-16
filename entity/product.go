package entity

type Product struct {
	Uuid           uint32  `json:"uuid"`
	CategoryId     uint32  `json:"categoryId"`
	ColorId        uint32  `json:"colorId"`
	SeasonId       uint32  `json:"seasonId"`
	SizeId         uint32  `json:"sizeId"`
	ManufacturerId uint32  `json:"manufacturerId"`
	BrandId        uint32  `json:"brandId"`
	GenderId       uint32  `json:"genderId"`
	AgeGroupId     uint32  `json:"ageGroupId"`
	PriceRoubles   float64 `json:"priceRoubles"`
}
