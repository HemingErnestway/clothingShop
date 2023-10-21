package storage

import (
	"clothingShop/dto"
	"clothingShop/entity"
	"reflect"
	"sync"
)

type CartMx struct {
	mtx   sync.RWMutex
	iter  uint32
	carts map[uint32]entity.Cart
}

var cartMx CartMx

func init() {
	cartMx = CartMx{
		carts: make(map[uint32]entity.Cart),
	}
}

func CartCreate(cart entity.Cart) *entity.Cart {
	cartMx.mtx.Lock()
	defer cartMx.mtx.Unlock()

	cartMx.iter++
	cart.Uuid = cartMx.iter
	cartMx.carts[cartMx.iter] = cart

	return &cart
}

func CartRead(id uint32) *entity.Cart {
	cartMx.mtx.RLock()
	defer cartMx.mtx.RUnlock()

	if el, ok := cartMx.carts[id]; ok {
		return &el
	}

	return nil
}

func CartUpdate(new dto.Cart, id uint32) *entity.Cart {
	cartMx.mtx.Lock()
	defer cartMx.mtx.Unlock()

	current := cartMx.carts[id]

	if new.TotalPrice != 0.0 {
		current.TotalPrice = new.TotalPrice
	}

	for _, attr := range new.ClearAttr {
		s := reflect.ValueOf(&current).Elem()
		field := s.FieldByName(attr)
		if field.CanSet() {
			field.SetZero()
		}
	}

	cartMx.carts[id] = current
	return &current
}

func CartDelete(id uint32) string {
	cartMx.mtx.Lock()
	defer cartMx.mtx.Unlock()

	delete(cartMx.carts, id)

	return "successfully deleted"
}
