package dto

import "clothingShop/entity"

type Cart struct {
	entity.Cart
	ClearAttr []string
}
