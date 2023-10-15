package dto

import "clothingShop/entity"

type User struct {
	entity.User
	ClearAttr []string `json:"clearAttr"`
}
