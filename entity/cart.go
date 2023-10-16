package entity

// type CartItem struct {
// 	ProductId        uint32  `json:"productId"`
// 	Quantity         uint32  `json:"quantity"`
// 	ItemPriceRoubles float64 `json:"itemPriceRoubles"`
// }
//
// type Cart struct {
// 	Items             []CartItem `json:"items"`
// 	TotalPriceRoubles float64    `json:"totalPriceRoubles"`

type CartItem struct {
	Uuid      uint32 `json:"uuid"`
	ProductId uint32 `json:"productId"`
	Quantity  uint32 `json:"quantity"`
	ItemPrice uint32 `json:"itemPrice"`
}

type Cart struct {
	Uuid       uint32   `json:"uuid"`
	UserId     uint32   `json:"userId"`
	TotalPrice float64  `json:"totalPrice"`
	ItemIds    []uint32 `json:"itemIds"`
}
