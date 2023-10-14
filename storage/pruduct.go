package storage

import (
	"clothingShop/entity"
	"sync"
)

type ProductMx struct {
	mtx      sync.RWMutex
	iter     uint32
	products map[uint32]entity.Product
}

var productMx ProductMx

func init() {
	productMx = ProductMx{
		products: make(map[uint32]entity.Product),
	}
}

func ProductCreate(product entity.Product) *entity.Product {
	productMx.mtx.Lock()
	defer productMx.mtx.Lock()
	productMx.iter++
	product.Uuid = productMx.iter
	productMx.products[productMx.iter] = product
	return &product
}

func ProductGetAll() []entity.Product {
	productMx.mtx.RLocker()
	defer productMx.mtx.RUnlock()
	lst := make([]entity.Product, len(productMx.products))
	iter := 0
	for key := range productMx.products {
		lst[iter] = productMx.products[key]
		iter++
	}
	return lst
}

func ProductGet(uid uint32) *entity.Product {
	productMx.mtx.RLock()
	defer productMx.mtx.RUnlock()
	if el, ok := productMx.products[uid]; ok {
		return &el
	}
	return nil
}
