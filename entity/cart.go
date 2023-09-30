package entity

import "github.com/google/uuid"

type CartItem struct {
	ProductId        uuid.UUID
	Quantity         uint32
	ItemPriceRoubles float64
}

type Cart struct {
	Items             []CartItem
	TotalPriceRoubles float64
}
