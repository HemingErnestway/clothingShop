package dto

import "clothingShop/entity"

type Product struct {
	entity.Product
	ClearAttr []string `json:"clearAttr"`
}
