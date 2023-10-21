package entity

type CartItem struct {
	Uuid      uint32 `json:"uuid"`
	CartId    uint32 `json:"cartId"`
	ProductId uint32 `json:"productId"`
	Quantity  uint32 `json:"quantity"`
	ItemPrice uint32 `json:"itemPrice"`
}

type Cart struct {
	Uuid       uint32  `json:"uuid"`
	UserId     uint32  `json:"userId"`
	TotalPrice float64 `json:"totalPrice"`
}
