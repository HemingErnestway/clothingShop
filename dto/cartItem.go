package dto

import "clothingShop/entity"

type CartItem struct {
	entity.CartItem
	ClearAttr []string `json:"clearAttr"`
}
