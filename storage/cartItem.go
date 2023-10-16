package storage

import (
	"clothingShop/entity"
	"sync"
)

type CartItemMx struct {
	mtx       sync.RWMutex
	iter      uint32
	cartItems map[uint32]entity.CartItem
}

var cartItemMx CartItemMx

func init() {
	cartItemMx = CartItemMx{
		cartItems: make(map[uint32]entity.CartItem),
	}
}
