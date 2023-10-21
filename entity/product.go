package entity

// type Product struct {
//	Uuid           uint32  `json:"uuid"`
//	CategoryId     uint32  `json:"categoryId"`
//	ColorId        uint32  `json:"colorId"`
//	SeasonId       uint32  `json:"seasonId"`
//	SizeId         uint32  `json:"sizeId"`
//	ManufacturerId uint32  `json:"manufacturerId"`
//	BrandId        uint32  `json:"brandId"`
//	GenderId       uint32  `json:"genderId"`
//	AgeGroupId     uint32  `json:"ageGroupId"`
//	PriceRoubles   float64 `json:"priceRoubles"`
// }

type Product struct {
	Uuid            uint32  `json:"uuid"`
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
	// TODO: implement size table
}
