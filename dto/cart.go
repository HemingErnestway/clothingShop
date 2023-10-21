package dto

import "clothingShop/entity"

type Cart struct {
	entity.Cart
	ClearAttr []string `json:"clearAttr"`
}

type CartItem struct {
	entity.CartItem
	ClearAttr []string `json:"clearAttr"`
}
